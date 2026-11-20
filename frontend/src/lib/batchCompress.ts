/** 批量压缩：扫描目录 -> 逐张压缩 -> 直接替换原图 */
import {
  ListImageFiles,
  ReadImageBase64,
  ReplaceImageBase64,
} from "../../wailsjs/go/main/App";
import { compressImage } from "./imageCompress";
import { base64ToBytes, bytesToBase64 } from "./bytes";

export interface BatchProgress {
  /** 待压缩文件总数 */
  total: number;
  /** 已处理数量 */
  done: number;
  /** 当前文件名 */
  current: string;
  /** 成功压缩并替换的数量 */
  compressed: number;
  /** 跳过的数量（未达标、变大或失败） */
  skipped: number;
  /** 未能压到目标大小的数量 */
  belowTarget: number;
  /** 处理失败的数量 */
  failed: number;
  /** 压缩前总字节数（仅成功替换的） */
  beforeBytes: number;
  /** 压缩后总字节数（仅成功替换的） */
  afterBytes: number;
}

function joinPath(folder: string, name: string): string {
  const base = folder.replace(/[\\/]+$/, "");
  return `${base}/${name}`;
}

/**
 * 压缩目录中超过目标大小的图片，并直接替换原图。
 * 只有“达到目标且比原图更小”的结果才会写回。
 */
export async function compressLargeImages(
  folder: string,
  targetBytes: number,
  onProgress: (progress: BatchProgress) => void
): Promise<BatchProgress> {
  const files = await ListImageFiles(folder);
  const state: BatchProgress = {
    total: files.length,
    done: 0,
    current: "",
    compressed: 0,
    skipped: 0,
    belowTarget: 0,
    failed: 0,
    beforeBytes: 0,
    afterBytes: 0,
  };
  if (files.length === 0) {
    return state;
  }

  for (const file of files) {
    state.current = file.name;
    onProgress({ ...state });

    const fullPath = joinPath(folder, file.name);
    try {
      const base64 = await ReadImageBase64(fullPath);
      const input = base64ToBytes(base64);
      const result = await compressImage(file.name, input, targetBytes);

      if (result.targetMet && result.data.byteLength < input.byteLength) {
        const outBase64 = await bytesToBase64(result.data);
        await ReplaceImageBase64(fullPath, outBase64);
        state.compressed++;
        state.beforeBytes += file.size;
        state.afterBytes += result.data.byteLength;
      } else {
        state.skipped++;
        if (!result.targetMet) {
          state.belowTarget++;
        }
      }
    } catch (err) {
      console.error(`[compress] ${file.name} 失败`, err);
      state.failed++;
      state.skipped++;
    }

    state.done++;
    onProgress({ ...state });
    // 让出一帧，保证进度条能刷新
    await new Promise((resolve) => setTimeout(resolve, 0));
  }

  return state;
}
