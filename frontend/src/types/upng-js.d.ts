declare module "upng-js" {
  interface UPNGImage {
    width: number;
    height: number;
    data: Uint8Array;
    frames?: unknown[];
    tabs?: unknown;
  }

  interface UPNGStatic {
    /** 编码 PNG；cnum > 0 时使用 pngquant 算法量化成索引色 */
    encode(
      bufs: ArrayBuffer[],
      width: number,
      height: number,
      cnum?: number,
      dels?: number[]
    ): ArrayBuffer;
    decode(buffer: ArrayBuffer): UPNGImage;
    toRGBA8(img: UPNGImage): ArrayBuffer[];
  }

  const UPNG: UPNGStatic;
  export default UPNG;
}
