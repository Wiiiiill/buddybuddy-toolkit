/**
 * 图片压缩核心逻辑。
 *
 * JPEG: MozJPEG（squoosh 同款 wasm），按质量阶梯编码，取第一个能进目标大小的质量。
 * PNG:  先用 UPNG 做无损重编码；仍超标时用 pngquant 算法（UPNG 的量化实现）
 *       量化成索引色 PNG。
 *
 * 说明：MozJPEG 与 pngquant 量化都是“视觉无损/感知无损”，并非像素级无损。
 */
import decodeJpeg from "@jsquash/jpeg/decode.js";
import encodeJpeg from "@jsquash/jpeg/encode.js";
import UPNG from "upng-js";

export type ImageKind = "jpeg" | "png";

export interface CompressResult {
  /** 压缩后的文件内容 */
  data: Uint8Array;
  /** 是否达到目标大小 */
  targetMet: boolean;
  /** 使用的策略，便于界面展示，例如 "MozJPEG q82" */
  method: string;
}

// 质量从高到低尝试。q90 通常已接近视觉无损；保留 58 作为兜底。
const JPEG_QUALITIES = [90, 82, 74, 66, 58];
// PNG 量化颜色数从多到少尝试。
const PNG_COLOR_STEPS = [256, 192, 128, 96, 64, 48, 32];

const yieldToUi = () => new Promise<void>((resolve) => setTimeout(resolve, 0));

/** 根据扩展名判断可处理的图片类型 */
export function imageKind(fileName: string): ImageKind | null {
  const dot = fileName.lastIndexOf(".");
  if (dot < 0) return null;
  const ext = fileName.slice(dot).toLowerCase();
  if (ext === ".jpg" || ext === ".jpeg") return "jpeg";
  if (ext === ".png") return "png";
  return null;
}

/** 把 Uint8Array 复制成独立 ArrayBuffer，避免 wasm 读取到多余字节 */
function exactBuffer(bytes: Uint8Array): ArrayBuffer {
  return bytes.buffer.slice(
    bytes.byteOffset,
    bytes.byteOffset + bytes.byteLength
  ) as ArrayBuffer;
}

/** 压缩一张图片，返回压缩结果；不负责写盘 */
export async function compressImage(
  fileName: string,
  input: Uint8Array,
  targetBytes: number
): Promise<CompressResult> {
  const kind = imageKind(fileName);
  if (kind === "jpeg") {
    return compressJpeg(exactBuffer(input), targetBytes);
  }
  if (kind === "png") {
    return compressPng(exactBuffer(input), targetBytes);
  }
  throw new Error(`不支持的文件类型: ${fileName}`);
}

async function compressJpeg(
  buffer: ArrayBuffer,
  targetBytes: number
): Promise<CompressResult> {
  // preserveOrientation: 让解码器按 EXIF 方向旋转像素，重编码后方向仍然正确
  const imageData = await decodeJpeg(buffer, { preserveOrientation: true });

  let last: Uint8Array | null = null;
  let lastQuality = JPEG_QUALITIES[0];
  for (const quality of JPEG_QUALITIES) {
    // 每轮让出主线程，保证进度可以刷新
    await yieldToUi();
    const encoded = new Uint8Array(
      await encodeJpeg(imageData, { quality })
    );
    last = encoded;
    lastQuality = quality;
    if (encoded.byteLength <= targetBytes) {
      return {
        data: encoded,
        targetMet: true,
        method: `MozJPEG q${quality}`,
      };
    }
  }

  return {
    data: last as Uint8Array,
    targetMet: false,
    method: `MozJPEG q${lastQuality}`,
  };
}

async function compressPng(
  buffer: ArrayBuffer,
  targetBytes: number
): Promise<CompressResult> {
  // 解码为 RGBA，再逐级量化成索引色 PNG
  const image = UPNG.decode(buffer);
  const rgba = new Uint8Array(UPNG.toRGBA8(image)[0]);

  // 第一步：纯无损重编码（cnum=0），像素完全不变
  const lossless = quantizePng(rgba, image.width, image.height, 0);
  if (lossless.byteLength <= targetBytes && lossless.byteLength < buffer.byteLength) {
    return { data: lossless, targetMet: true, method: "PNG 无损重编码" };
  }

  let last: Uint8Array | null = null;
  let lastMethod = "";
  for (const colors of PNG_COLOR_STEPS) {
    await yieldToUi();
    const encoded = quantizePng(rgba, image.width, image.height, colors);
    last = encoded;
    lastMethod = `量化 ${colors} 色`;
    if (encoded.byteLength <= targetBytes) {
      return { data: encoded, targetMet: true, method: lastMethod };
    }
  }

  return { data: last as Uint8Array, targetMet: false, method: lastMethod };
}

/** 用 UPNG（pngquant 算法移植）量化成索引色 PNG */
function quantizePng(
  rgba: Uint8Array,
  width: number,
  height: number,
  colors: number
): Uint8Array {
  const buffer = rgba.buffer.slice(
    rgba.byteOffset,
    rgba.byteOffset + rgba.byteLength
  ) as ArrayBuffer;
  const encoded = UPNG.encode([buffer], width, height, colors);
  return new Uint8Array(encoded);
}
