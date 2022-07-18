<script context="module">
  import { CodeJar } from "codejar";
  import { withLineNumbers } from "codejar/linenumbers";

  const Prism = window.Prism;
  export function codedit(
    node,
    { code, autofocus = false, disabled = false, loc = true, ...options }
  ) {
    const highlight = loc
      ? withLineNumbers(Prism.highlightElement)
      : Prism.highlightElement;
    if (disabled) {
      function update() {
        node.textContent = code;
        highlight(node);
      }
      update();
      return {
        update,
      };
    } else {
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
  }
</script>

<script>
  export let id = undefined;
  export let code = "";
  export let autofocus = false;
  export let loc = false;
  export let disabled = false;
  export let popover = false;
  export let maxHeight = false;
</script>

<pre
  {id}
  use:codedit={{ code, autofocus, loc, disabled, ...$$restProps }}
  readonly={disabled}
  class="language-sql"
  class:popover
  class:max-height={maxHeight}
  on:change
/>

<style>
  .popover {
    z-index: 1;
    position: relative;
  }

  .max-height {
    max-height: 500px;
    overflow-y: scroll;
  }
</style>
