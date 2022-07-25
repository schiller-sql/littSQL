import { get, writable, type Writable } from "svelte/store";

/// low means that nothing is done about performance, high means components are more optimized, but have less features
export type PerformanceMode = "low" | "high";

const localStorageKey = "performance_mode";

function getPerformanceModeFromLocalStorage(): PerformanceMode {
  return (
    (localStorage.getItem(localStorageKey) as PerformanceMode | null) ?? "low"
  );
}

export let performanceMode = getPerformanceModeFromLocalStorage();
export const performanceModeStore: Writable<PerformanceMode> =
  writable(performanceMode);

performanceModeStore.subscribe((performanceM) => {
  localStorage.setItem(localStorageKey, performanceM);
  performanceMode = performanceM;
});

export function togglePerformanceMode() {
  return get(performanceModeStore) === "low" ? "high" : "loq";
}
