import { defineStore } from "pinia";
import { GetConfig, SaveConfig } from "../../wailsjs/go/main/App";
import { main } from "../../wailsjs/go/models";

export interface Preset {
  name: string;
  before: string;
  after: string;
}

export const useConfigStore = defineStore("config", {
  state: () => ({
    supportedFormats: [] as string[],
    presets: [] as Preset[],
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
    },
    async save() {
      await SaveConfig(
        new main.Config({
          supportedFormats: this.supportedFormats,
          presets: this.presets,
        })
      );
    },
  },
});
