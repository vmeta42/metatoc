<template>
  <a-modal :visible="showPopup" @cancel="onHandleClose">
    <template #title>
      <div
        :style="{
          display: 'flex',
        }"
      >
        <div
          :style="{
            paddingTop: '6px',
          }"
        >
          <a-avatar :size="36" :src="handleUserAvatar(avatar)" />
        </div>
        <div
          :style="{
            marginLeft: '16px',
            width: '90%',
          }"
        >
          <div>
            <span
              :style="{
                fontSize: '12px',
                fontWeight: 'normal',
                color: 'rgba(0, 0, 0, 0.45)',
              }"
              ></span
            >
          </div>
          <span>{{  }}</span>
          <div>
          </div>
        </div>
      </div>
    </template>
    <div
      :style="{
        display: 'flex',
      }"
    >
      <div>
      </div>
      <div
        :style="{
          paddingLeft: '16px',
          // color: 'rgba(0, 0, 0, 0.45)',
          width: '90%',
          listStyleType: 'none',
          // margin: '10px auto'
        }"
      >
        <ul v-for="(item,$index) in newTxtList" :key="$index">
          <li
            :style="{
              listStyleType: 'none',
              // margin: '10px auto'
            }"
            >
            <span>文件名称</span>
            <a-button 
              style="cursor: 'pointer',"
              @click="downLoadfile(item.file_name)"
              type="link">
              {{ item.file_name.split("/")[2].toString() }}
            </a-button>
          </li>
        </ul>
        <!-- <chat-on-chain-view-dialog-markdown-parser
          :onChainChatData="onChainChatData"
        /> -->
      </div>
    </div>
    <template #footer>
      <a-button @click="onHandleClose">取消</a-button>
    </template>
  </a-modal>
</template>

<script>
import { ref, watchEffect } from "vue";
// import ChatOnChainViewDialogMarkdownParser from "./ChatOnChainViewDialogMarkdownParser.vue";
// import moment from "moment";
import EmilyJohnsonAvatar from "@/assets/avatar/pexels-photo-16187929.jpeg";
import MichaelSmithAvatar from "@/assets/avatar/pexels-photo-16161525.jpeg";
import SophiaWilliamsAvatar from "@/assets/avatar/pexels-photo-16196205.jpeg";
import { $post, $get } from "@/utils/request";

const usePopup = (props, emit) => {
  const showPopup = ref(false);
  const title = ref("");
  const updateAt = ref("");
  const avatar = ref("");

  const newAddress = ref("")
  const newPrivateKey = ref("")
  const newPath = ref("")
  // const newAvatar = ref("")
  const newTxtList = ref([])

  const getViewDetail = () => {
    if(!newAddress.value || !newPrivateKey.value || !newPath.value ) {
      return false
    }
    $post("/metatoc-service/v1/blockchain/view", {
      address: newAddress.value,
      private_key: newPrivateKey.value,
      path: newPath.value,
    }).then((res) => {
      $post("/metatoc-service/v1/metadata/listFiles", {
        object_uuid: res.data.data
      }).then((res) => {
        newTxtList.value = res.data.data
      })
    })
  }

  watchEffect(() => {
    newAddress.value = props.nowViewData.address
    newPrivateKey.value = props.nowViewData.privateKey
    newPath.value = props.nowViewData.path
    // newAvatar.value = props.nowViewData.avatar
    showPopup.value = props.viewDataVisible || false

    getViewDetail()
    // title.value = props.onChainChatData.chat || "";
    // avatar.value =
    //   (props.onChainChatData.avatarSrc && props.onChainChatData.avatarSrc[0]) ||
    //   "";
    // updateAt.value = moment(Number(props.onChainChatData.updateAt)).format(
    //   "YYYY-MM-DD HH:mm:SS"
    // );
  });
  const onHandleClose = () => {
    emit("popupClose");
  };
  return {
    // getViewDetail,
    newTxtList,


    showPopup,
    title,
    avatar,
    updateAt,
    onHandleClose,
  };
};

export default {
  name: "ChatOnChainViewDialog",
  props: {
    viewDataVisible: {
      type: Boolean,
      default: false,
    },
    newCurrentUser: {
      type: Object,
    },
    nowViewData: {
      type: Object,
    }
  },
  components: {
    // ChatOnChainViewDialogMarkdownParser,
  },
  methods: {
    handleUserAvatar() {
      return EmilyJohnsonAvatar
      // if (avatar == "@/assets/avatar/pexels-photo-16187929.jpeg") {
      //   return EmilyJohnsonAvatar;
      // } else if (avatar == "@/assets/avatar/pexels-photo-16161525.jpeg") {
      //   return MichaelSmithAvatar;
      // } else if (avatar == "@/assets/avatar/pexels-photo-16196205.jpeg") {
      //   return SophiaWilliamsAvatar;
      // }
    },
  },
  data() {
    return {
      EmilyJohnsonAvatar,
      MichaelSmithAvatar,
      SophiaWilliamsAvatar,
    };
  },
  setup(props, { emit }) {
    const { showPopup, title, avatar, updateAt, onHandleClose, newTxtList } = usePopup(
      props,
      emit
    );

    const downLoadfile = (item) => {
      $get("/metatoc-service/v1/metadata/download", {
        "object_name": item
      }).then((res) => {
        const url = window.URL.createObjectURL(new Blob([res.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', item);
        document.body.appendChild(link);
        link.click();
      })
    }
   


    console.log(props.nowViewData);
    return {
      showPopup,
      title,
      avatar,
      updateAt,
      onHandleClose,
      newTxtList,
      downLoadfile
    };
  },
};
</script>

<style scoped></style>
