export interface ImageFile {
  id: number;
  base64: string;
  mimeType: string;
  path: string;
}

export interface ConvertedImage {
  id: number;
  name: string;
  format: string;
  size: number;
  blob: Blob;
  url: string;
  originalImage: ImageFile;
}

export interface Toast {
  id: number;
  message: string;
  type: "success" | "error" | "warning" | "info";
  timestamp: number;
}
