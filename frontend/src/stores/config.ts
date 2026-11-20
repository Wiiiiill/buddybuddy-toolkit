import { defineStore } from "pinia";
import { GetConfig, SaveConfig, ResetConfig } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

export interface Preset {
  name: string;
  before: string;
  after: string;
}

// daisyUI 内置主题名
export const THEMES = [
  "light", "dark", "cupcake", "bumblebee", "emerald", "corporate",
  "synthwave", "retro", "cyberpunk", "valentine", "halloween", "garden",
  "forest", "aqua", "lofi", "pastel", "fantasy", "wireframe", "black",
  "luxury", "dracula", "cmyk", "autumn", "business", "acid", "lemonade",
  "night", "coffee", "winter", "dim", "nord", "sunset",
];

export const useConfigStore = defineStore("config", {
  state: () => ({
    supportedFormats: [] as string[],
    presets: [] as Preset[],
    locale: "zh-Hans",
    theme: "light",
  }),
  actions: {
    async load() {
      const cfg = await GetConfig();
      this.supportedFormats = cfg.supportedFormats ?? [];
      this.presets = (cfg.presets ?? []).map((p) => ({
        name: p.name,
        before: p.before,
        after: p.after,
      }));
      this.locale = cfg.locale || "zh-Hans";
      this.theme = cfg.theme || "light";
      this.applyTheme();
    },
    async save() {
      await SaveConfig(
        new main.Config({
          supportedFormats: this.supportedFormats,
          presets: this.presets,
          locale: this.locale,
          theme: this.theme,
        })
      );
    },
    applyTheme() {
      document.documentElement.setAttribute("data-theme", this.theme);
    },
    async setTheme(theme: string) {
      this.theme = theme;
      this.applyTheme();
      await this.save();
    },
    async setLocale(locale: string) {
      this.locale = locale;
      await this.save();
    },
    async reset() {
      const cfg = await ResetConfig();
      this.supportedFormats = cfg.supportedFormats ?? [];
      this.presets = (cfg.presets ?? []).map((p) => ({
        name: p.name,
        before: p.before,
        after: p.after,
      }));
      this.locale = cfg.locale || "zh-Hans";
      this.theme = cfg.theme || "light";
      this.applyTheme();
    },
  },
});
