import { defineStore } from "pinia";

type Theme = "light" | "dark";

export const useThemeStore = defineStore("theme", {
  state: () => ({
    theme: (localStorage.getItem("theme") as Theme) || "light",
  }),
  actions: {
    init() {
      this.apply();
    },
    toggle() {
      this.theme = this.theme === "light" ? "dark" : "light";
      localStorage.setItem("theme", this.theme);
      this.apply();
    },
    apply() {
      document.documentElement.setAttribute("data-theme", this.theme);
    },
  },
});
