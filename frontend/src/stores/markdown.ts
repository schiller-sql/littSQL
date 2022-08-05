import { writable, type Readable } from "svelte/store";
import { marked } from "marked";
import DOMPurify from "dompurify";

// TODO: use web workers

export type MarkdownRenderStatus =
  | {
      status: "loading";
    }
  | { status: "rendered"; renderedMarkdown: string };

interface MarkdownRenderStore extends Readable<MarkdownRenderStatus> {
  rawMarkdownUpdate(rawMarkdown: string): void;
}

const domPurifyConfig = {
  USE_PROFILES: { html: true },
};

export function createMarkdownRenderStore(
  millisecondsTillCheck: number | undefined
): MarkdownRenderStore {
  let lastRawMarkdown: string | undefined;
  const w = writable<MarkdownRenderStatus>({ status: "loading" });
  let lastTimeoutId;
  function renderRawMarkdown(rawMarkdown: string) {
    const dirtyRenderedMarkdown = marked(rawMarkdown);
    const cleanRenderedMarkdown = DOMPurify.sanitize(
      dirtyRenderedMarkdown,
      domPurifyConfig
    );
    w.set({ status: "rendered", renderedMarkdown: cleanRenderedMarkdown });
  }
  function rawMarkdownUpdate(rawMarkdown: string) {
    if (lastRawMarkdown === rawMarkdown) return;
    lastRawMarkdown = rawMarkdown;
    if (millisecondsTillCheck !== undefined) {
      w.set({ status: "loading" });
      clearTimeout(lastTimeoutId);
      lastTimeoutId = setTimeout(
        renderRawMarkdown,
        millisecondsTillCheck,
        rawMarkdown
      );
    } else {
      renderRawMarkdown(rawMarkdown);
    }
  }
  return {
    rawMarkdownUpdate,
    subscribe: w.subscribe,
  };
}
