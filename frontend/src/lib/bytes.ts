/** base64 与二进制互转（用于跨越 Wails 绑定传输文件内容） */

/** base64 -> Uint8Array */
export function base64ToBytes(base64: string): Uint8Array {
  const binary = atob(base64);
  const out = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    out[i] = binary.charCodeAt(i);
  }
  return out;
}

/** Uint8Array -> base64，借助 FileReader 处理大文件，避免 btoa 参数过长 */
export async function bytesToBase64(bytes: Uint8Array): Promise<string> {
  const blob = new Blob([bytes]);
  return await new Promise<string>((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      const result = String(reader.result);
      resolve(result.slice(result.indexOf(",") + 1));
    };
    reader.onerror = () => reject(reader.error);
    reader.readAsDataURL(blob);
  });
}
