import { reactive } from "vue";

export const toast = reactive({
  message: "",
  type: "success" as "success" | "error",
  visible: false,
});

let timer: ReturnType<typeof setTimeout> | null = null;

export function showToast(message: string, type: "success" | "error" = "success") {
  toast.message = message;
  toast.type = type;
  toast.visible = true;
  if (timer) clearTimeout(timer);
  timer = setTimeout(() => {
    toast.visible = false;
  }, 2500);
}
