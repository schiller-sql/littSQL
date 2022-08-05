import { writable, type Readable } from "svelte/store";

/// low means that nothing is done about performance, high means components are more optimized, but have less features
export type PerformanceMode = "low" | "high";

const localStorageKey = "performance_mode";

function getPerformanceModeFromLocalStorage(): PerformanceMode {
  return (
    (localStorage.getItem(localStorageKey) as PerformanceMode | null) ?? "low"
  );
}

interface PerformanceStore extends Readable<PerformanceMode> {
  togglePerformanceMode(): void;
  getCurrentMode(): PerformanceMode;
}

export function createPerformanceStore(): PerformanceStore {
  let currentMode = getPerformanceModeFromLocalStorage();
  const w = writable<PerformanceMode>(currentMode);
  w.subscribe((performanceM) => {
    localStorage.setItem(localStorageKey, performanceM);
    currentMode = performanceM;
  });
  return {
    subscribe: w.subscribe,
    togglePerformanceMode(): void {
      w.set(currentMode === "low" ? "high" : "low");
    },
    getCurrentMode(): PerformanceMode {
      return currentMode;
    },
  };
}
