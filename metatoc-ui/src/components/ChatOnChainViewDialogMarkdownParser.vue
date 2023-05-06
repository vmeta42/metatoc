<template>
  <div ref="markdownParser"></div>
</template>

<script>
import { ref } from "vue";
import showdown from "showdown";

export default {
  name: "ChatOnChainViewDialogMarkdownParser",
  props: {
    onChainChatData: {
      type: Object,
    },
  },
  setup(props) {
    const converter = new showdown.Converter();

    const innerHTML = ref("");
    innerHTML.value = props.onChainChatData.reply || "";

    return {
      converter,
      innerHTML,
    };
  },
  watch: {
    onChainChatData(newValue) {
      this.innerHTML = newValue.reply || "";
      const html = this.converter.makeHtml(this.innerHTML);
      this.$refs.markdownParser.innerHTML = html;
    },
  },
  mounted() {
    this.$nextTick(() => {
      const html = this.converter.makeHtml(this.innerHTML);
      this.$refs.markdownParser.innerHTML = html;
    });
  },
};
</script>

<style scoped></style>
