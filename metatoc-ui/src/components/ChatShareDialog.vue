<template>
  <a-popover trigger="click" v-model:visible="visible">
    <template #title>
      <span>Share</span>
    </template>
    <template #content>
      <p>Please select who you want to shared</p>
      <a-select
        v-model:value="value"
        :options="buildOptions(userList, item.users)"
        mode="multiple"
        placeholder="Please select"
        style="width: 230px"
        @dropdownVisibleChange="dropdownVisibleChange"
      >
      </a-select>
      <div
        :style="{
          display: 'flex',
          justifyContent: 'end',
          marginTop: '12px',
        }"
      >
        <a-button :size="size" @click="cancel">Cancel</a-button>
        <a-button
          type="primary"
          :size="size"
          :style="{
            marginLeft: '6px',
          }"
          @click="ok(this.item)"
          v-if="nodeStatus == 'Stopped'"
          disabled
          >OK</a-button
        >
        <a-button
          type="primary"
          :size="size"
          :style="{
            marginLeft: '6px',
          }"
          @click="ok(this.item)"
          v-else-if="nodeStatus == 'Running'"
          :loading="loadingPopup"
          >OK</a-button
        >
      </div>
    </template>
    <share-alt-outlined key="share" />
  </a-popover>
</template>

<script>
import { ref } from "vue";
import { ShareAltOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { share } from "../utils/block";

export default {
  name: "ChatShareDialog",
  props: {
    item: {
      type: Object,
    },
    userList: {
      type: Array,
    },
    nodeStatus: {
      type: String,
    },
  },
  watch: {
    item(newValue) {
      console.log("newValue====>", newValue);
    },
  },
  components: {
    ShareAltOutlined,
  },
  methods: {
    buildOptions(userList, users) {
      console.log("buildOptions");
      console.log("users====>", users);
      const options = [];
      userList.forEach((element) => {
        if (users.indexOf(element) == -1) {
          options.push({
            value: element,
          });
        }
      });
      return options;
    },
  },
  setup(props) {
    const dropdownVisibleChange = (open) => {
      console.log("call func [dropdownVisibleChange]");
      console.log("open", open);
    };

    const value = ref([]);
    const options = ref([]);
    props.userList.forEach((element) => {
      options.value.push({
        value: element,
      });
    });

    const visible = ref(false);
    const loadingPopup = ref(false);
    const ok = async (item) => {
      console.log(item);
      if (value.value.length > 0) {
        let successNum = 0;
        loadingPopup.value = true;
        const private_key = JSON.parse(localStorage.getItem("currentUser"))[
          "private_key"
        ];
        const from_address = JSON.parse(localStorage.getItem("currentUser"))[
          "address"
        ];
        for (const element of value.value) {
          let to_address = "";
          JSON.parse(localStorage.getItem("userList")).forEach((user) => {
            if (user.name == element) {
              to_address = user.address;
            }
          });
          let token_name = "/" + item.uuid;

          try {
            const res = await share(
              private_key,
              from_address,
              to_address,
              token_name
            );
            console.log("res==>", res);
            if (res && res.code == 0 && res.message == "SUCCESSFUL") {
              item.users.push(element);
              const onChainChatArray = JSON.parse(
                localStorage.getItem("onChainChat")
              );
              onChainChatArray.forEach((element) => {
                if (element.uuid == item.uuid) {
                  element.users = item.users;
                  element.updateAt = Date.now();
                }
              });
              localStorage.setItem(
                "onChainChat",
                JSON.stringify(onChainChatArray)
              );
              successNum++;
            }
          } catch (err) {
            console.log("err==>", err);
          }
        }
        if (successNum == value.value.length) {
          message.success("Share chat successful!");
        } else {
          message.error("Share chat failed, please try again!");
        }
        visible.value = false;
        loadingPopup.value = false;
        value.value = [];
      } else {
        message.error("Please select who you want to shared!");
      }
    };

    const cancel = () => {
      visible.value = false;
      value.value = [];
    };

    const size = ref("small");

    return {
      dropdownVisibleChange,
      // value: ref(["a1", "b2"]),
      // options: [...Array(25)].map((_, i) => ({
      //   value: (i + 10).toString(36) + (i + 1),
      // })),
      value,
      options,
      visible,
      loadingPopup,
      ok,
      cancel,
      size,
    };
  },
};
</script>

<style scoped></style>
