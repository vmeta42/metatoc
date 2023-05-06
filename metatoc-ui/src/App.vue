<template>
  <a-layout>
    <a-layout-header class="header">
      <div class="logo">
        <img
          src="../src/assets/logo.png"
          alt=""
          :style="{
            width: '40px',
          }"
        />
        <span
          :style="{
            color: '#fff',
            fontWeight: 'bolder',
            fontSize: '18px',
            marginLeft: '12px',
            cursor: 'pointer',
          }"
          @click="redirctToGithub"
          >MetaTOC</span
        >
      </div>
      <div class="avatar">
        <!-- <a-avatar
          :size="32"
          :style="{ backgroundColor: color, verticalAlign: 'middle' }"
        >
          {{ avatarValue }}
        </a-avatar> -->
        <!-- <p>{{ JSON.parse(localCurrentUser)["avatar"] }}</p> -->
        <a-avatar
          :size="36"
          :src="handleUserAvatar(JSON.parse(localCurrentUser)['avatar'])"
          :style="{ cursor: 'pointer' }"
          @click="onOpenSwitchUserPopup"
        />
        <caret-down-outlined
          :style="{
            color: '#eee',
            paddingLeft: '6px',
            cursor: 'pointer',
          }"
          @click="onOpenSwitchUserPopup"
        />
        <!-- <a-button
          size="small"
          :style="{ marginLeft: '16px', verticalAlign: 'middle' }"
          @click="changeValue"
        >
          改变
        </a-button> -->
      </div>
    </a-layout-header>
    <a-layout>
      <a-layout-sider width="200" style="background: #fff">
        <a-card
          class="hovered"
          :bordered="false"
          style="width: 200px"
          :bodyStyle="cardBodyStyle"
          @click="onOpenPopup()"
        >
          <plus-outlined />
          <span :style="{ paddingLeft: '14px' }">New chat</span>
        </a-card>
        <a-divider style="margin: 0px 0px" />
        <a-card
          v-for="item in renderChatArray"
          :key="item.uuid"
          class="hovered"
          :bordered="false"
          style="width: 200px"
          :bodyStyle="cardBodyStyle"
          @click="onOpenPopup(item)"
        >
          <cloud-outlined v-if="item.state == 'off-chain'" />
          <cloud-upload-outlined v-else />
          <span :style="{ paddingLeft: '14px' }">{{ item.chat }}</span>
        </a-card>
      </a-layout-sider>
      <a-layout style="padding: 0 0 0 24px">
        <a-layout-content
          :style="{
            background: '#fff',
            margin: 0,
            minHeight: '280px',
          }"
        >
          <a-card :bordered="false" :bodyStyle="cardBodyStyle">
            <block-outlined />
            <span :style="{ paddingLeft: '14px' }">On chain chat</span>
          </a-card>

          <a-divider style="margin: 0px 0px" />

          <div style="padding: 24px">
            <ul class="nav">
              <li
                v-for="(item, index) in renderOnChainChatArray"
                :key="item.uuid"
              >
                <!-- <a-card style="width: 300px"> -->
                <a-card :style="handleLiCardStyle(item, index)">
                  <template #title>
                    <a-card-meta>
                      <template #title>
                        <div>
                          <span
                            :style="{
                              fontSize: '12px',
                              fontWeight: 'normal',
                              color: 'rgba(0, 0, 0, 0.45)',
                            }"
                            >{{
                              moment(Number(item.updateAt)).format(
                                "YYYY-MM-DD HH:mm:ss"
                              )
                            }}</span
                          >
                        </div>
                        <span>{{ item.chat }}</span>
                      </template>
                      <template #avatar>
                        <div
                          :style="{
                            paddingTop: '6px',
                          }"
                        >
                          <!-- <a-avatar
                            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                          /> -->
                          <a-avatar
                            :size="36"
                            :src="handleUserAvatar(item.avatarSrc[0])"
                          />
                        </div>
                      </template>
                    </a-card-meta>
                  </template>
                  <a-card-meta>
                    <template #description>
                      <div
                        style="
                          height: 175px;
                          display: -webkit-box;
                          word-break: break-all;
                          text-overflow: ellipsis;
                          -webkit-box-orient: vertical;
                          -webkit-line-clamp: 8;
                          overflow: hidden;
                        "
                      >
                        <span>{{ item.reply }}</span>
                      </div>
                    </template>
                    <template #avatar>
                      <svg
                        width="36"
                        height="36"
                        viewBox="-3 0 46 46"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                        strokewidth="2"
                        class="scale-appear"
                      >
                        <path
                          d="M37.5324 16.8707C37.9808 15.5241 38.1363 14.0974 37.9886 12.6859C37.8409 11.2744 37.3934 9.91076 36.676 8.68622C35.6126 6.83404 33.9882 5.3676 32.0373 4.4985C30.0864 3.62941 27.9098 3.40259 25.8215 3.85078C24.8796 2.7893 23.7219 1.94125 22.4257 1.36341C21.1295 0.785575 19.7249 0.491269 18.3058 0.500197C16.1708 0.495044 14.0893 1.16803 12.3614 2.42214C10.6335 3.67624 9.34853 5.44666 8.6917 7.47815C7.30085 7.76286 5.98686 8.3414 4.8377 9.17505C3.68854 10.0087 2.73073 11.0782 2.02839 12.312C0.956464 14.1591 0.498905 16.2988 0.721698 18.4228C0.944492 20.5467 1.83612 22.5449 3.268 24.1293C2.81966 25.4759 2.66413 26.9026 2.81182 28.3141C2.95951 29.7256 3.40701 31.0892 4.12437 32.3138C5.18791 34.1659 6.8123 35.6322 8.76321 36.5013C10.7141 37.3704 12.8907 37.5973 14.9789 37.1492C15.9208 38.2107 17.0786 39.0587 18.3747 39.6366C19.6709 40.2144 21.0755 40.5087 22.4946 40.4998C24.6307 40.5054 26.7133 39.8321 28.4418 38.5772C30.1704 37.3223 31.4556 35.5506 32.1119 33.5179C33.5027 33.2332 34.8167 32.6547 35.9659 31.821C37.115 30.9874 38.0728 29.9178 38.7752 28.684C39.8458 26.8371 40.3023 24.6979 40.0789 22.5748C39.8556 20.4517 38.9639 18.4544 37.5324 16.8707ZM22.4978 37.8849C20.7443 37.8874 19.0459 37.2733 17.6994 36.1501C17.7601 36.117 17.8666 36.0586 17.936 36.0161L25.9004 31.4156C26.1003 31.3019 26.2663 31.137 26.3813 30.9378C26.4964 30.7386 26.5563 30.5124 26.5549 30.2825V19.0542L29.9213 20.998C29.9389 21.0068 29.9541 21.0198 29.9656 21.0359C29.977 21.052 29.9842 21.0707 29.9867 21.0902V30.3889C29.9842 32.375 29.1946 34.2791 27.7909 35.6841C26.3872 37.0892 24.4838 37.8806 22.4978 37.8849ZM6.39227 31.0064C5.51397 29.4888 5.19742 27.7107 5.49804 25.9832C5.55718 26.0187 5.66048 26.0818 5.73461 26.1244L13.699 30.7248C13.8975 30.8408 14.1233 30.902 14.3532 30.902C14.583 30.902 14.8088 30.8408 15.0073 30.7248L24.731 25.1103V28.9979C24.7321 29.0177 24.7283 29.0376 24.7199 29.0556C24.7115 29.0736 24.6988 29.0893 24.6829 29.1012L16.6317 33.7497C14.9096 34.7416 12.8643 35.0097 10.9447 34.4954C9.02506 33.9811 7.38785 32.7263 6.39227 31.0064ZM4.29707 13.6194C5.17156 12.0998 6.55279 10.9364 8.19885 10.3327C8.19885 10.4013 8.19491 10.5228 8.19491 10.6071V19.808C8.19351 20.0378 8.25334 20.2638 8.36823 20.4629C8.48312 20.6619 8.64893 20.8267 8.84863 20.9404L18.5723 26.5542L15.206 28.4979C15.1894 28.5089 15.1703 28.5155 15.1505 28.5173C15.1307 28.5191 15.1107 28.516 15.0924 28.5082L7.04046 23.8557C5.32135 22.8601 4.06716 21.2235 3.55289 19.3046C3.03862 17.3858 3.30624 15.3413 4.29707 13.6194ZM31.955 20.0556L22.2312 14.4411L25.5976 12.4981C25.6142 12.4872 25.6333 12.4805 25.6531 12.4787C25.6729 12.4769 25.6928 12.4801 25.7111 12.4879L33.7631 17.1364C34.9967 17.849 36.0017 18.8982 36.6606 20.1613C37.3194 21.4244 37.6047 22.849 37.4832 24.2684C37.3617 25.6878 36.8382 27.0432 35.9743 28.1759C35.1103 29.3086 33.9415 30.1717 32.6047 30.6641C32.6047 30.5947 32.6047 30.4733 32.6047 30.3889V21.188C32.6066 20.9586 32.5474 20.7328 32.4332 20.5338C32.319 20.3348 32.154 20.1698 31.955 20.0556ZM35.3055 15.0128C35.2464 14.9765 35.1431 14.9142 35.069 14.8717L27.1045 10.2712C26.906 10.1554 26.6803 10.0943 26.4504 10.0943C26.2206 10.0943 25.9948 10.1554 25.7963 10.2712L16.0726 15.8858V11.9982C16.0715 11.9783 16.0753 11.9585 16.0837 11.9405C16.0921 11.9225 16.1048 11.9068 16.1207 11.8949L24.1719 7.25025C25.4053 6.53903 26.8158 6.19376 28.2383 6.25482C29.6608 6.31589 31.0364 6.78077 32.2044 7.59508C33.3723 8.40939 34.2842 9.53945 34.8334 10.8531C35.3826 12.1667 35.5464 13.6095 35.3055 15.0128ZM14.2424 21.9419L10.8752 19.9981C10.8576 19.9893 10.8423 19.9763 10.8309 19.9602C10.8195 19.9441 10.8122 19.9254 10.8098 19.9058V10.6071C10.8107 9.18295 11.2173 7.78848 11.9819 6.58696C12.7466 5.38544 13.8377 4.42659 15.1275 3.82264C16.4173 3.21869 17.8524 2.99464 19.2649 3.1767C20.6775 3.35876 22.0089 3.93941 23.1034 4.85067C23.0427 4.88379 22.937 4.94215 22.8668 4.98473L14.9024 9.58517C14.7025 9.69878 14.5366 9.86356 14.4215 10.0626C14.3065 10.2616 14.2466 10.4877 14.2479 10.7175L14.2424 21.9419ZM16.071 17.9991L20.4018 15.4978L24.7325 17.9975V22.9985L20.4018 25.4983L16.071 22.9985V17.9991Z"
                          fill="currentColor"
                        ></path>
                      </svg>
                    </template>
                  </a-card-meta>
                  <template #actions>
                    <expand-alt-outlined
                      key="expand"
                      @click="onOpenViewPopup(item)"
                    />
                    <chat-share-dialog
                      :item="item"
                      :userList="userList"
                      :nodeStatus="nodeStatus"
                      @shareSubmit="shareSubmit"
                    ></chat-share-dialog>
                  </template>
                </a-card>
              </li>
            </ul>
          </div>

          <a-divider style="margin: 0px 0px" />

          <a-card :bordered="false" :bodyStyle="cardBodyStyle">
            <deployment-unit-outlined />
            <span :style="{ paddingLeft: '14px' }">Node status</span>
            <span :style="{ color: '#1890ff' }" v-if="nodeStatus == 'Running'"
              >&nbsp;(💙Running
            </span>
            <span
              class="heartbeat-text"
              :style="{ color: '#1890ff' }"
              v-if="nodeStatus == 'Running'"
            >
              Beng~ Beng~
            </span>
            <span :style="{ color: '#1890ff' }" v-if="nodeStatus == 'Running'"
              >)
            </span>
            <!-- <div class="heartbeat-text"></div> -->
            <span
              :style="{ color: '#f5222d' }"
              v-else-if="nodeStatus == 'Stopped'"
              >&nbsp;(💔Stopped)
            </span>
            <!-- <div id="container" /> -->
          </a-card>

          <a-divider style="margin: 0px 0px" />

          <a-card :bordered="false" :bodyStyle="cardBodyStyle">
            <node-controller-vue
              @nodeStatusChange="onNodeStatusChange"
            ></node-controller-vue>
          </a-card>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>

  <chat-create-and-edit-dialog
    :visible="visible"
    :loading="loading"
    :isEdit="isEdit"
    :chatData="chatData"
    :nodeStatus="nodeStatus"
    @popupClose="onClosePopup"
    @popupDelete="onDeletePopup"
    @popupSubmit="onSubmitPopup"
    @popupToChain="onToChainPopup"
  />
  <chat-on-chain-view-dialog
    :visible="chatViewDialogVisible"
    :onChainChatData="onChainChatData"
    @popupClose="onCloseViewPopup"
  />
  <switch-user-dialog
    :visible="switchUserDialogVisible"
    :userList="userList"
    :avatarSrc="avatarSrc"
    :localCurrentUser="localCurrentUser"
    :localUserList="localUserList"
    @popupClose="onCloseSwitchUserPopup"
    @popupSubmit="onSubmitSwitchUserPopup"
  />
  <full-loading-vue :fullLoading="fullLoading"></full-loading-vue>
</template>
<script>
import { ref, defineComponent, watchEffect } from "vue";
import {
  PlusOutlined,
  CloudOutlined,
  BlockOutlined,
  DeploymentUnitOutlined,
  ExpandAltOutlined,
  // ShareAltOutlined,
  CloudUploadOutlined,
  CaretDownOutlined,
} from "@ant-design/icons-vue";
import ChatCreateAndEditDialog from "./components/ChatCreateAndEditDialog.vue";
import ChatOnChainViewDialog from "./components/ChatOnChainViewDialog.vue";
import ChatShareDialog from "./components/ChatShareDialog.vue";
import SwitchUserDialog from "./components/SwitchUserDialog.vue";
import NodeControllerVue from "./components/NodeController.vue";
import { v4 as uuidv4 } from "uuid";
import { $post } from "./utils/request";
import moment from "moment";
import EmilyJohnsonAvatar from "@/assets/avatar/pexels-photo-16187929.jpeg";
import MichaelSmithAvatar from "@/assets/avatar/pexels-photo-16161525.jpeg";
import SophiaWilliamsAvatar from "@/assets/avatar/pexels-photo-16196205.jpeg";
import { message } from "ant-design-vue";
import { initChatString, initOnChainChatString } from "./utils/contents";
import FullLoadingVue from "./components/FullLoading.vue";
import { dataInitialization } from "./utils/init";
import { create } from "./utils/block";

const usePopup = (props, emit, { chatArray, onChainChatArray }) => {
  const visible = ref(false);
  const loading = ref(false);
  const isEdit = ref(false);
  const chatData = ref({});
  const onOpenPopup = (thisChatData) => {
    if (thisChatData == undefined) {
      // New chat
      visible.value = true;
      loading.value = false;
      isEdit.value = false;
      chatData.value = {};
    } else {
      // Edit chat
      visible.value = true;
      loading.value = false;
      isEdit.value = true;
      chatData.value = thisChatData;
    }
  };
  const onClosePopup = () => {
    visible.value = false;
    loading.value = false;
  };
  const onDeletePopup = ({ chatUuid, chatValue }) => {
    console.log("call func [onDeletePopup]");
    console.log("chatUuid==>", chatUuid);
    console.log("chatValue==>", chatValue);
    const newChatArray = [];
    chatArray.value.reverse();
    chatArray.value.forEach((element) => {
      if (element.uuid != chatUuid.value) {
        newChatArray.push(element);
      }
    });
    console.log("newChatArray==>", newChatArray);
    console.log("newChatArray.reverse()==>", newChatArray);
    localStorage.setItem("chat", JSON.stringify(newChatArray));
    newChatArray.reverse();
    chatArray.value = newChatArray;
    visible.value = false;
    loading.value = false;
  };
  const onSubmitPopup = ({ chatUuid, chatValue }) => {
    const newChatArray = [];
    if (chatArray.value.length == 0) {
      const chatObject = {
        uuid: uuidv4(),
        chat: chatValue.value,
        reply: "",
        state: "off-chain",
        users: [JSON.parse(localStorage.getItem("currentUser"))["name"]],
        createAt: Date.now(),
        updateAt: Date.now(),
      };
      newChatArray.push(chatObject);
    } else {
      chatArray.value.reverse();
      let inChatArray = false;
      chatArray.value.forEach((element) => {
        if (element.uuid == chatUuid.value) {
          inChatArray = true;
          element.chat = chatValue.value;
          element.updateAt = Date.now();
        }
        newChatArray.push(element);
      });
      if (inChatArray == false) {
        const chatObject = {
          uuid: uuidv4(),
          chat: chatValue.value,
          reply: "",
          state: "off-chain",
          users: [JSON.parse(localStorage.getItem("currentUser"))["name"]],
          createAt: Date.now(),
          updateAt: Date.now(),
        };
        newChatArray.push(chatObject);
      }
    }
    localStorage.setItem("chat", JSON.stringify(newChatArray));
    newChatArray.reverse();
    chatArray.value = newChatArray;
    visible.value = false;
    loading.value = false;
  };

  const onToChainPopup = ({ chatUuid, chatValue }) => {
    $post("/sendMsg", { msg: chatValue.value }).then(
      (res) => {
        let chatObject = {};
        if (chatUuid.value == "") {
          chatObject = {
            uuid: uuidv4(),
            chat: chatValue.value,
            reply: res.data,
            state: "on-chain",
            users: [JSON.parse(localStorage.getItem("currentUser"))["name"]],
            avatarSrc: [
              JSON.parse(localStorage.getItem("currentUser"))["avatar"],
            ],
            createAt: Date.now(),
            updateAt: Date.now(),
          };
        } else {
          chatArray.value.forEach((element) => {
            if (element.uuid == chatUuid.value) {
              chatObject = {
                uuid: element.uuid,
                chat: element.chat,
                reply: res.data,
                state: "on-chain",
                users: [
                  JSON.parse(localStorage.getItem("currentUser"))["name"],
                ],
                avatarSrc: [
                  JSON.parse(localStorage.getItem("currentUser"))["avatar"],
                ],
                createAt: element.createAt,
                updateAt: Date.now(),
              };
            }
          });
        }
        console.log("chatObject==>", chatObject);
        console.log("chatObject.length==>", Object.keys(chatObject).length);
        if (Object.keys(chatObject).length != 0) {
          create(
            JSON.parse(localStorage.getItem("currentUser"))["private_key"],
            JSON.parse(localStorage.getItem("currentUser"))["address"],
            "/" + chatObject["uuid"],
            JSON.stringify({
              chat: chatObject["chat"],
              reply: chatObject["reply"],
            })
          ).then(
            (res) => {
              console.log(res);
              if (chatUuid.value == "") {
                chatArray.value.unshift(chatObject);
                onChainChatArray.value.unshift(chatObject);
              } else {
                chatArray.value.forEach((element) => {
                  if (element.uuid == chatUuid.value) {
                    element.reply = chatObject["reply"];
                    element.state = chatObject["state"];
                    element.users = chatObject["users"];
                    element.avatarSrc = chatObject["avatarSrc"];
                    element.updateAt = chatObject["updateAt"];
                  }
                });
                onChainChatArray.value.unshift(chatObject);
              }
              visible.value = false;
              loading.value = false;
              const localChatArray = chatArray.value;
              const localOnChainChatArray = onChainChatArray.value;
              localStorage.setItem(
                "chat",
                JSON.stringify(localChatArray.slice().reverse())
              );
              localStorage.setItem(
                "onChainChat",
                JSON.stringify(localOnChainChatArray.slice().reverse())
              );
              message.success("Chat successfully added to the blockchain!");
            },
            (err) => {
              console.log(err);
              visible.value = true;
              loading.value = false;
              message.success("Chat failed to be added to the blockchain!");
            }
          );
        }
      },
      (err) => {
        console.log("err==>", err);
        let chatObject = {};
        if (chatUuid.value == "") {
          chatObject = {
            uuid: uuidv4(),
            chat: chatValue.value,
            reply:
              "Unable to connect to ChatGPT service, please check your network connection or try again later.",
            state: "on-chain",
            users: [JSON.parse(localStorage.getItem("currentUser"))["name"]],
            avatarSrc: [
              JSON.parse(localStorage.getItem("currentUser"))["avatar"],
            ],
            createAt: Date.now(),
            updateAt: Date.now(),
          };
        } else {
          chatArray.value.forEach((element) => {
            if (element.uuid == chatUuid.value) {
              chatObject = {
                uuid: element.uuid,
                chat: element.chat,
                reply: element.reply
                  ? element.reply
                  : "Unable to connect to ChatGPT service, please check your network connection or try again later.",
                state: "on-chain",
                users: [
                  JSON.parse(localStorage.getItem("currentUser"))["name"],
                ],
                avatarSrc: [
                  JSON.parse(localStorage.getItem("currentUser"))["avatar"],
                ],
                createAt: element.createAt,
                updateAt: Date.now(),
              };
            }
          });
        }
        console.log("chatObject==>", chatObject);
        console.log("chatObject.length==>", Object.keys(chatObject).length);
        if (Object.keys(chatObject).length != 0) {
          create(
            JSON.parse(localStorage.getItem("currentUser"))["private_key"],
            JSON.parse(localStorage.getItem("currentUser"))["address"],
            "/" + chatObject["uuid"],
            JSON.stringify({
              chat: chatObject["chat"],
              reply: chatObject["reply"],
            })
          ).then(
            (res) => {
              console.log(res);
              if (chatUuid.value == "") {
                chatArray.value.unshift(chatObject);
                onChainChatArray.value.unshift(chatObject);
              } else {
                chatArray.value.forEach((element) => {
                  if (element.uuid == chatUuid.value) {
                    element.reply = chatObject["reply"];
                    element.state = chatObject["state"];
                    element.users = chatObject["users"];
                    element.avatarSrc = chatObject["avatarSrc"];
                    element.updateAt = chatObject["updateAt"];
                  }
                });
                onChainChatArray.value.unshift(chatObject);
              }
              visible.value = false;
              loading.value = false;
              const localChatArray = chatArray.value;
              const localOnChainChatArray = onChainChatArray.value;
              localStorage.setItem(
                "chat",
                JSON.stringify(localChatArray.slice().reverse())
              );
              localStorage.setItem(
                "onChainChat",
                JSON.stringify(localOnChainChatArray.slice().reverse())
              );
              message.success("Chat successfully added to the blockchain!");
            },
            (err) => {
              console.log(err);
              visible.value = true;
              loading.value = false;
              message.success("Chat failed to be added to the blockchain!");
            }
          );
        }
      }
    );
  };

  const onOpenSharePopup = (thisChatData) => {
    console.log(thisChatData);
  };
  return {
    visible,
    loading,
    isEdit,
    chatData,
    onOpenPopup,
    onClosePopup,
    onDeletePopup,
    onSubmitPopup,
    onToChainPopup,
    onOpenSharePopup,
  };
};

const useViewPopup = () => {
  const chatViewDialogVisible = ref(false);
  const onChainChatData = ref({});
  const onOpenViewPopup = (thisOnChainChatData) => {
    chatViewDialogVisible.value = true;
    onChainChatData.value = thisOnChainChatData;
  };
  const onCloseViewPopup = () => {
    chatViewDialogVisible.value = false;
  };
  return {
    chatViewDialogVisible,
    onChainChatData,
    onOpenViewPopup,
    onCloseViewPopup,
  };
};

const switchUserPopup = ({ userList, avatarSrc }) => {
  const switchUserDialogVisible = ref(false);

  let myUsers = [];
  userList.forEach((element, index) => {
    if (index == 0) {
      myUsers.push({
        id: index,
        name: element,
        avatar: avatarSrc[index],
        lastLogin: moment().format("YYYY-MM-DD HH:mm:ss"),
      });
    } else {
      myUsers.push({
        id: index,
        name: element,
        avatar: avatarSrc[index],
        lastLogin: moment().format("YYYY-MM-DD HH:mm:ss"),
      });
    }
  });
  const localCurrentUser = ref("");
  localCurrentUser.value = localStorage.getItem("currentUser") || "";
  if (localCurrentUser.value == "") {
    localCurrentUser.value = JSON.stringify(myUsers[0]);
    localStorage.setItem("currentUser", localCurrentUser.value);
  }
  const localUserList = ref([]);
  localUserList.value = localStorage.getItem("userList") || "";
  console.log("localUserList.value.length==>", localUserList.value.length);
  if (localUserList.value.length == 0) {
    localUserList.value = JSON.stringify(myUsers);
    localStorage.setItem("userList", localUserList.value);
  }

  const onOpenSwitchUserPopup = () => {
    switchUserDialogVisible.value = true;
  };
  const onCloseSwitchUserPopup = () => {
    switchUserDialogVisible.value = false;
  };
  const onSubmitSwitchUserPopup = (selectedUserId) => {
    console.log("selectedUserId==>", selectedUserId);
    let localCurrentUserObject = {};
    let localUserListArray = JSON.parse(localStorage.getItem("userList"));
    localUserListArray.forEach((element) => {
      if (selectedUserId == element["id"]) {
        localCurrentUserObject = element;
        element["lastLogin"] = moment().format("YYYY-MM-DD HH:mm:ss");
      }
    });
    localCurrentUser.value = JSON.stringify(localCurrentUserObject);
    localStorage.setItem("currentUser", localCurrentUser.value);
    localUserList.value = JSON.stringify(localUserListArray);
    localStorage.setItem("userList", localUserList.value);
    switchUserDialogVisible.value = false;
  };
  return {
    switchUserDialogVisible,
    localCurrentUser,
    localUserList,
    onOpenSwitchUserPopup,
    onCloseSwitchUserPopup,
    onSubmitSwitchUserPopup,
  };
};

const nodeStatusChangeClass = () => {
  const nodeStatus = ref("Running");
  const onNodeStatusChange = (thisNodeStatus) => {
    nodeStatus.value = thisNodeStatus;
  };
  return {
    nodeStatus,
    onNodeStatusChange,
  };
};

export default defineComponent({
  components: {
    PlusOutlined,
    CloudOutlined,
    BlockOutlined,
    DeploymentUnitOutlined,
    ExpandAltOutlined,
    // ShareAltOutlined,
    CloudUploadOutlined,
    CaretDownOutlined,
    ChatCreateAndEditDialog,
    ChatOnChainViewDialog,
    ChatShareDialog,
    SwitchUserDialog,
    NodeControllerVue,
    FullLoadingVue,
  },
  methods: {
    redirctToGithub() {
      window.open("https://github.com/vmeta42/metatoc");
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
  data() {
    return {
      EmilyJohnsonAvatar,
      MichaelSmithAvatar,
      SophiaWilliamsAvatar,
    };
  },
  setup(props, emit) {
    const fullLoading = ref(true);

    const userList = ["Emily Johnson", "Michael Smith", "Sophia Williams"];
    const avatarSrc = [
      "@/assets/avatar/pexels-photo-16187929.jpeg",
      "@/assets/avatar/pexels-photo-16161525.jpeg",
      "@/assets/avatar/pexels-photo-16196205.jpeg",
    ];

    // 初始化chat
    if (!localStorage.getItem("chat")) {
      // const initChat = [
      //   {
      //     avatarSrc: [avatarSrc[0]],
      //     chat: "What is Fermat's Last Theorem?",
      //     createAt: Date.now(),
      //     reply: `Fermat's Last Theorem, proposed by Pierre de Fermat in 1637, states that no three positive integers a, b, and c satisfy the equation an + bn = cn for any integer value of n greater than 2. This theorem was famously proved by Andrew Wiles in 1994, after more than 350 years of effort by mathematicians to solve it.`,
      //     state: "on-chain",
      //     updateAt: Date.now(),
      //     users: [userList[0]],
      //     uuid: uuidv4(),
      //   },
      //   {
      //     avatarSrc: [avatarSrc[1]],
      //     chat: "What is the largest star in the universe?",
      //     createAt: Date.now(),
      //     reply:
      //       "The largest star currently known is UY Scuti, which has a radius around 1,700 times larger than the Sun. However, there may be even larger stars that have not been discovered yet.",
      //     state: "on-chain",
      //     updateAt: Date.now(),
      //     users: [userList[1]],
      //     uuid: uuidv4(),
      //   },
      //   {
      //     avatarSrc: [avatarSrc[2]],
      //     chat: "Can you use Go to write a code that collects CPU and memory?",
      //     createAt: Date.now(),
      //     reply:
      //       'Yes, it\'s possible to use Go language to collect CPU and memory usage of a program or system.\n\nGo provides standard library packages such as "runtime" and "os" which can be used to gather information about system resources such as CPU usage, memory usage, and more.\n\nFor example, you can use the "runtime" package to get information about the current Go process:\n\n```go\nimport "runtime"\n\nfunc main() {\n    memStats := &runtime.MemStats{}\n    runtime.ReadMemStats(memStats)\n    // memStats now contains the memory usage statistics for the current process\n}\n```\n\nSimilarly, you can use the "os" package to get information about the overall system:\n\n```go\nimport (\n    "fmt"\n    "os"\n)\n\nfunc main() {\n    stat := &syscall.Statfs_t{}\n\n    err := syscall.Statfs("/", stat)\n    if err != nil {\n        fmt.Println("Error getting file system stats:", err)\n        return\n    }\n\n    totalSpace := stat.Blocks * uint64(stat.Bsize)\n    freeSpace := stat.Bfree * uint64(stat.Bsize)\n\n    // totalSpace and freeSpace now contain the total and free disk space on the root file system\n}\n```\n\nThese are just examples, and there are many other ways to gather system resource usage data using Go.',
      //     state: "on-chain",
      //     updateAt: Date.now(),
      //     users: [userList[2]],
      //     uuid: uuidv4(),
      //   },
      // ];
      localStorage.setItem("chatUnconfirm", initChatString);
      localStorage.setItem("onChainChatUnconfirm", initOnChainChatString);
    }

    const chat = ref("");
    const chatArray = ref([]);
    chat.value = localStorage.getItem("chat") || "";
    if (chat.value != "") {
      chatArray.value = JSON.parse(chat.value);
      chatArray.value.reverse();
    }

    const onChainChat = ref("");
    const onChainChatArray = ref([]);
    onChainChat.value = localStorage.getItem("onChainChat") || "";
    if (onChainChat.value != "") {
      onChainChatArray.value = JSON.parse(onChainChat.value);
      onChainChatArray.value.reverse();
    }

    const {
      visible,
      loading,
      isEdit,
      chatData,
      onOpenPopup,
      onClosePopup,
      onDeletePopup,
      onSubmitPopup,
      onToChainPopup,
      onOpenSharePopup,
    } = usePopup(props, emit, { chatArray, onChainChatArray });

    const {
      chatViewDialogVisible,
      onChainChatData,
      onOpenViewPopup,
      onCloseViewPopup,
    } = useViewPopup();

    const cardBodyStyle = ref({
      padding: "16px 24px",
      overflow: "hidden",
      textOverflow: "ellipsis",
      whiteSpace: "nowrap",
      wordBreak: "break-all",
      width: "100%",
    });
    const colorList = ["#f56a00", "#f56a10", "#f56a20"];
    const avatarValue = ref(userList[0]);
    const color = ref(colorList[0]);
    const changeValue = () => {
      const index = userList.indexOf(avatarValue.value);
      avatarValue.value =
        index < userList.length - 1 ? userList[index + 1] : userList[0];
      color.value =
        index < colorList.length - 1 ? colorList[index + 1] : colorList[0];
    };

    const handleLiCardStyle = (item, index) => {
      console.log("item==>", item);
      console.log("index==>", index);
      if (index == 0) {
        return {
          width: "300px",
        };
      } else {
        return {
          width: "300px",
          marginLeft: "24px",
        };
      }
    };

    const shareSubmit = (item) => {
      console.log("in [shareSubmit]");
      console.log("item==>", item);
      const onChainChatArray = JSON.parse(localStorage.getItem("onChainChat"));
      onChainChatArray.forEach((element) => {
        if (element.uuid == item.uuid) {
          element.users = item.users;
          element.updateAt = Date.now();
        }
      });
      localStorage.setItem("onChainChat", JSON.stringify(onChainChatArray));
      console.log("onChainChatArray==>", onChainChatArray);
    };

    const {
      switchUserDialogVisible,
      localCurrentUser,
      localUserList,
      onOpenSwitchUserPopup,
      onCloseSwitchUserPopup,
      onSubmitSwitchUserPopup,
    } = switchUserPopup({ userList, avatarSrc });

    const renderChatArray = ref([]);
    const renderOnChainChatArray = ref([]);
    watchEffect(() => {
      renderChatArray.value = [];
      renderOnChainChatArray.value = [];
      console.log("localCurrentUser.value==>", localCurrentUser.value);
      chatArray.value.forEach((element) => {
        if (
          element.users.indexOf(JSON.parse(localCurrentUser.value)["name"]) > -1
        ) {
          renderChatArray.value.push(element);
        }
      });
      onChainChatArray.value.forEach((element) => {
        if (
          element.users.indexOf(JSON.parse(localCurrentUser.value)["name"]) > -1
        ) {
          renderOnChainChatArray.value.push(element);
        }
      });
    });

    const { onNodeStatusChange, nodeStatus } = nodeStatusChangeClass();

    if (localStorage.getItem("initState") != "Finish") {
      dataInitialization();
      const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
      const checkInitState = async () => {
        let initState = localStorage.getItem("initState");
        while (initState !== "Finish") {
          chat.value = localStorage.getItem("chat") || "";
          if (chat.value != "") {
            chatArray.value = JSON.parse(chat.value);
            chatArray.value.reverse();
          }

          onChainChat.value = localStorage.getItem("onChainChat") || "";
          if (onChainChat.value != "") {
            onChainChatArray.value = JSON.parse(onChainChat.value);
            onChainChatArray.value.reverse();
          }

          console.log("initState:", initState);
          await delay(1000);
          initState = localStorage.getItem("initState");
        }
      };
      (async () => {
        await checkInitState();
        chat.value = localStorage.getItem("chat") || "";
        if (chat.value != "") {
          chatArray.value = JSON.parse(chat.value);
          chatArray.value.reverse();
        }

        onChainChat.value = localStorage.getItem("onChainChat") || "";
        if (onChainChat.value != "") {
          onChainChatArray.value = JSON.parse(onChainChat.value);
          onChainChatArray.value.reverse();
        }
        fullLoading.value = false;
      })();
    } else {
      fullLoading.value = false;
    }

    return {
      chatArray,
      onChainChatArray,
      renderChatArray,
      renderOnChainChatArray,

      visible,
      loading,
      isEdit,
      chatData,
      onOpenPopup,
      onClosePopup,
      onDeletePopup,
      onSubmitPopup,
      onToChainPopup,
      onOpenSharePopup,

      chatViewDialogVisible,
      onChainChatData,
      onOpenViewPopup,
      onCloseViewPopup,

      userList,
      avatarSrc,
      cardBodyStyle,
      avatarValue,
      color,
      changeValue,

      handleLiCardStyle,

      moment,

      shareSubmit,

      switchUserDialogVisible,
      localCurrentUser,
      localUserList,
      onOpenSwitchUserPopup,
      onCloseSwitchUserPopup,
      onSubmitSwitchUserPopup,

      nodeStatus,
      onNodeStatusChange,

      fullLoading,
    };
  },
});
</script>

<style scoped>
/* #components-layout-demo-top-side-2 .logo {
  float: left;
  width: 120px;
  height: 31px;
  margin: 16px 24px 16px 0;
  background: rgba(255, 255, 255, 0.3);
}
.ant-row-rtl #components-layout-demo-top-side-2 .logo {
  float: right;
  margin: 16px 0 16px 24px;
} */
.site-layout-background {
  background: #fff;
}
.logo {
  display: inline-flex;
  justify-content: end; /* 水平居右 */
  align-items: center; /* 垂直居中 */
  height: 100%;
  float: left;
}
.avatar {
  display: inline-flex;
  justify-content: end; /* 水平居右 */
  align-items: center; /* 垂直居中 */
  height: 100%;
  float: right;
}
.hovered:hover {
  background: #f5f5f5;
  cursor: pointer;
}

.nav {
  display: flex;
  /* padding-top: 24px; */
  /* padding-bottom: 24px; */
  margin: 0px;
  padding: 0px;
  /* padding-bottom: 12px; */
  /* background-color: #f3f5f7; */
  overflow-x: auto;
}
.nav li {
  /* flex: 1; */
  display: flex;
  /* flex-direction: column;
  font-size: 16px;
  justify-content: center;
  align-items: center; */
  /* width: 88px; */
  /* flex-basis: 88px; */
  /* flex-shrink: 0; */
  /* white-space: nowrap; */
}

.heartbeat-text {
  display: inline-block;
  animation: heartbeat 1.2s infinite;
}

@keyframes heartbeat {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  10% {
    transform: scale(1.1);
  }
  20% {
    transform: scale(1);
  }
  30% {
    transform: scale(1.1);
  }
  40% {
    transform: scale(1);
  }
  50% {
    transform: scale(0.9);
    opacity: 0.7;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}
</style>
