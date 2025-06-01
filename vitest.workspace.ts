import { defineWorkspace } from "vitest/config";

export default defineWorkspace([
  "./apps/website/vite.config.ts",
  "./apps/webapp/vite.config.ts",
  "./libs/elements/vite.config.ts",
  "./libs/provar-js/vite.config.ts",
  "./libs/mockgen/vite.config.ts",
  "./apps/collect/vite.config.ts",
  "./libs/types/vite.config.ts",
]);
