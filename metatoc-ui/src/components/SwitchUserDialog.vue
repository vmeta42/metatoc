<template>
  <div>
    <a-modal
      :visible="showPopup"
      title="Switch user"
      @cancel="onHandleClose"
      :destroyOnClose="true"
    >
      <div class="user-switch-container">
        <a-row>
          <a-col :span="24" v-for="user in myUsers" :key="user.id">
            <div
              class="user-card"
              :class="{ 'user-card-selected': user.id === selectedUserId }"
              @click="selectUser(user)"
            >
              <div class="user-avatar">
                <a-avatar :src="handleUserAvatar(user.avatar)" />
              </div>
              <div class="user-info">
                <div class="user-name">{{ user.name }}</div>
              </div>
              <div class="user-last-login" v-if="user.lastLogin != ''">
                Last login {{ user.lastLogin }}
              </div>
            </div>
          </a-col>
        </a-row>
      </div>
      <template #footer>
        <a-button @click="onHandleClose">Cancel</a-button>
        <a-button
          type="primary"
          @click="onHandleSubmit(this.selectedUserId)"
          :loading="loadingPopup"
          >OK</a-button
        >
      </template>
    </a-modal>
  </div>
</template>
<script>
import { defineComponent, ref, watchEffect } from "vue";
import { message } from "ant-design-vue";
import EmilyJohnsonAvatar from "@/assets/avatar/pexels-photo-16187929.jpeg";
import MichaelSmithAvatar from "@/assets/avatar/pexels-photo-16161525.jpeg";
import SophiaWilliamsAvatar from "@/assets/avatar/pexels-photo-16196205.jpeg";

const usePopup = (props, emit) => {
  const showPopup = ref(false);
  const loadingPopup = ref(false);
  const myUsers = ref([]);
  const selectedUserId = ref(0);
  watchEffect(() => {
    showPopup.value = props.visible || false;
    // props.userList.forEach((element, index) => {
    //   console.log("element==>", element);
    //   console.log("index==>", index);
    //   myUsers.value.push({
    //     id: index,
    //     name: element,
    //     avatar: props.avatarSrc[index],
    //     lastLogin: "2023-04-12 14:30",
    //   });
    // });
    myUsers.value = JSON.parse(props.localUserList);
    selectedUserId.value = JSON.parse(props.localCurrentUser)["id"];
  });
  const onHandleClose = () => {
    if (loadingPopup.value == false) {
      emit("popupClose");
    }
  };
  const onHandleSubmit = (selectedUserId) => {
    loadingPopup.value = true;
    setTimeout(() => {
      loadingPopup.value = false;
      emit("popupSubmit", selectedUserId);
      // message.success(
      //   "Switched user successfully! The current user is: " +
      //     JSON.parse(props.localUserList)[selectedUserId]["name"]
      // );
      message.success("Switch user successfully!");
    }, 0);
  };
  return {
    showPopup,
    loadingPopup,
    myUsers,
    selectedUserId,
    onHandleClose,
    onHandleSubmit,
  };
};

export default defineComponent({
  name: "SwitchUserDialog",
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    userList: {
      type: Array,
    },
    avatarSrc: {
      type: Array,
    },
    localCurrentUser: {
      type: String,
    },
    localUserList: {
      type: String,
    },
  },
  data() {
    return {
      EmilyJohnsonAvatar,
      MichaelSmithAvatar,
      SophiaWilliamsAvatar,
      // selectedUserId: null,
      // users: [
      //   {
      //     id: 1,
      //     name: "张三",
      //     email: "zhangsan@example.com",
      //     avatar: "https://example.com/avatar1.jpg",
      //     lastLogin: "2023-04-12 14:30",
      //   },
      //   {
      //     id: 2,
      //     name: "李四",
      //     email: "lisi@example.com",
      //     avatar: "https://example.com/avatar2.jpg",
      //     lastLogin: "2023-04-10 10:20",
      //   },
      //   // ...
      // ],
    };
  },
  methods: {
    selectUser(user) {
      this.selectedUserId = user.id;
    },
    handleUserAvatar(avatar) {
      if (avatar == "@/assets/avatar/pexels-photo-16187929.jpeg") {
        return EmilyJohnsonAvatar;
      } else if (avatar == "@/assets/avatar/pexels-photo-16161525.jpeg") {
        return MichaelSmithAvatar;
      } else if (avatar == "@/assets/avatar/pexels-photo-16196205.jpeg") {
        return SophiaWilliamsAvatar;
      }
    },
  },
  setup(props, { emit }) {
    const {
      showPopup,
      loadingPopup,
      onHandleClose,
      onHandleSubmit,
      myUsers,
      selectedUserId,
    } = usePopup(props, emit);

    return {
      showPopup,
      loadingPopup,
      onHandleClose,
      onHandleSubmit,
      myUsers,
      selectedUserId,
    };
  },
});
</script>

<style scoped>
.user-switch-container {
  max-height: 400px;
  overflow-y: auto;
}

.user-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  margin-bottom: 16px;
  transition: 0.3s;
  border: 2px solid transparent; /* 为未选中的卡片添加透明边框 */
}

.user-card-selected {
  border: 2px solid #1890ff;
}

.user-avatar {
  margin-right: 8px; /* 将头像和名字靠近 */
}

.user-info {
  display: flex; /* 让名字和头像靠近 */
  align-items: center;
  flex-grow: 1; /* 确保上次登录时间显示在卡片的最右边 */
}

.user-name {
  font-weight: 600;
  margin-left: 8px;
  margin-right: 8px; /* 让名字和头像紧挨着 */
}

.user-last-login {
  color: rgba(0, 0, 0, 0.45);
}
</style>
