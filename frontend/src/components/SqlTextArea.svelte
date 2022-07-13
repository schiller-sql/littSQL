<script context="module">
  import { CodeJar } from "codejar";
  import { withLineNumbers } from "codejar/linenumbers";
  import { onMount } from "svelte";
  const Prism = window.Prism;
  export function codedit(
    node,
    { code, autofocus = false, loc = true, ...options }
  ) {
    const highlight = loc
      ? withLineNumbers(Prism.highlightElement)
      : Prism.highlightElement;
    const editor = CodeJar(
      node,
      (n) => {
        n.textContent = n.textContent;
        return highlight(n);
      },
      options
    );
    editor.onUpdate((code) => {
      const e = new CustomEvent("change", { detail: code });
      node.dispatchEvent(e);
    });
    function update({ code, autofocus = false, loc = false, ...options }) {
      editor.updateOptions(options);
      if (editor.toString() !== code) {
        editor.updateCode(code);
      }
    }
    update({ code, autofocus, loc, ...options });
    autofocus && node.focus();
    return {
      update,
      destroy() {
        editor.destroy();
      },
    };
  }
</script>

<script>
  export let code = "";
  export let autofocus = false;
  export let loc = false;
  const style = "";
</script>

<pre
  use:codedit={{ code, autofocus, loc, ...$$restProps }}
  class="language-sql"
  {style}
  on:change
/>
