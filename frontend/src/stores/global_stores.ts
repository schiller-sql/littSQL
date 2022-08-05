import { createAuthStore } from "./auth";
import { createDatabaseStore } from "./database";
import { createPerformanceStore } from "./performance";

export const authStore = createAuthStore();
export const databaseStore = createDatabaseStore();
export const performanceStore = createPerformanceStore();
