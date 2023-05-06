<template>
  <div class="node-display">
    <div class="circle-container">
      <a-dropdown
        v-for="(node, index) in nodes"
        :key="index"
        :overlay="getNodeMenu(node, index)"
      >
        <div
          class="circle-node"
          :class="{ offline: node.status === 'offline' }"
          :style="nodePosition(index)"
          @mousedown="handleMousedown($event, index, node)"
        >
          NODE{{ index + 1 }}
        </div>
      </a-dropdown>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { Dropdown, Menu, Modal } from "ant-design-vue";
import { message } from "ant-design-vue";

export default defineComponent({
  name: "NodeDisplay",
  components: {
    [Dropdown.name]: Dropdown,
    [Menu.name]: Menu,
    [Menu.Item.name]: Menu.Item,
  },
  setup(props, { emit }) {
    const nodes = ref([]);
    const offlineNodeNumber = ref(0);
    const nodeState = ref("Running");
    if (localStorage.getItem("nodes")) {
      nodes.value = JSON.parse(localStorage.getItem("nodes"));
      offlineNodeNumber.value = localStorage.getItem("offlineNodeNumber");
      nodeState.value = localStorage.getItem("nodeState");
      emit("nodeStatusChange", nodeState.value);
    } else {
      nodes.value = reactive(
        Array.from({ length: 5 }, () => ({
          status: "online",
          messageStatus: "online",
          class: "circle-node",
        }))
      );
      localStorage.setItem("nodes", JSON.stringify(nodes.value));
      localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
      localStorage.setItem("nodeState", nodeState.value);
      emit("nodeStatusChange", nodeState.value);
    }

    const handleMousedown = (event, index) => {
      if (event.button === 0) {
        // 左键点击
        console.log(`节点${index + 1}被点击`);
      }
    };

    const nodePosition = (index) => {
      const radius = 120;
      const angle = -90 + (360 / nodes.value.length) * index;
      const x = radius * Math.cos((angle * Math.PI) / 180);
      const y = radius * Math.sin((angle * Math.PI) / 180);

      return {
        transform: `translate(-50%, -50%) translate(${x}px, ${y}px)`,
        fontSize: "12px",
      };
    };

    const getNodeMenu = (node, index) => {
      const items = [];
      if (node.messageStatus !== "offline") {
        items.push(<a-menu-item key="offline">Offline Node</a-menu-item>);
      }
      if (node.messageStatus !== "online") {
        items.push(<a-menu-item key="online">Online Node</a-menu-item>);
      }

      return (
        <a-menu
          onClick={({ key }) => {
            handleNodeStatusChange(index, node, key);
          }}
        >
          {items}
        </a-menu>
      );
    };

    const handleNodeStatusChange = (index, node, status) => {
      console.log("node==>", index);
      console.log("node==>", node);
      console.log("node==>", status);
      if (status == "offline") {
        if (offlineNodeNumber.value >= 2 && nodeState.value == "Running") {
          offlineNodeNumber.value++;
          localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
          showConfirmForOffline(
            index,
            node,
            status,
            offlineNodeNumber,
            nodeState
          );
        } else {
          offlineNodeNumber.value++;
          localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
          node.status = status;
          setTimeout(() => handleSetTimeout(index, node, status), 300);

          let newNodes = JSON.parse(localStorage.getItem("nodes"));
          newNodes[index] = node;
          localStorage.setItem("nodes", JSON.stringify(newNodes));
          message.success("Offline node successfully!");
        }
      } else {
        if (offlineNodeNumber.value <= 3 && nodeState.value == "Stopped") {
          offlineNodeNumber.value--;
          localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
          showConfirmForOnline(
            index,
            node,
            status,
            offlineNodeNumber,
            nodeState
          );
        } else {
          offlineNodeNumber.value--;
          localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
          node.status = status;
          setTimeout(() => handleSetTimeout(index, node, status), 300);

          let newNodes = JSON.parse(localStorage.getItem("nodes"));
          newNodes[index] = node;
          localStorage.setItem("nodes", JSON.stringify(newNodes));
          message.success("Onlinee node successfully!");
        }
      }
    };

    const handleSetTimeout = (index, node, status) => {
      node.messageStatus = status;
      let newNodes = JSON.parse(localStorage.getItem("nodes"));
      newNodes[index] = node;
      localStorage.setItem("nodes", JSON.stringify(newNodes));
    };

    const showConfirmForOnline = (
      index,
      node,
      status,
      offlineNodeNumber,
      nodeState
    ) => {
      Modal.info({
        title: "Node is about to go online.",
        content: `After this node goes online, the number of offline nodes will be below the threshold for the normal operation of the blockchain, and functions such as data on-chain and on-chain data sharing will become available again.`,
        onOk() {
          node.status = status;
          nodeState.value = "Running";
          setTimeout(() => handleSetTimeout(index, node, status), 300);

          let newNodes = JSON.parse(localStorage.getItem("nodes"));
          newNodes[index] = node;
          localStorage.setItem("nodes", JSON.stringify(newNodes));
          localStorage.setItem("nodeState", nodeState.value);

          emit("nodeStatusChange", nodeState.value);
          message.success("Online node successfully!");
        },
      });
    };

    const showConfirmForOffline = (
      index,
      node,
      status,
      offlineNodeNumber,
      nodeState
    ) => {
      Modal.confirm({
        title: "Do you want to take this node offline?",
        content: `After this node goes offline, the number of offline nodes will exceed the threshold for the normal operation of the blockchain, which will result in functions such as data on-chain and on-chain data sharing becoming unavailable.`,
        onOk() {
          node.status = status;
          nodeState.value = "Stopped";
          setTimeout(() => handleSetTimeout(index, node, status), 300);

          let newNodes = JSON.parse(localStorage.getItem("nodes"));
          newNodes[index] = node;
          localStorage.setItem("nodes", JSON.stringify(newNodes));
          localStorage.setItem("nodeState", nodeState.value);

          emit("nodeStatusChange", nodeState.value);
          message.success("Offline node successfully!");
        },
        onCancel() {
          offlineNodeNumber.value--;
          localStorage.setItem("offlineNodeNumber", offlineNodeNumber.value);
        },
      });
    };

    return {
      nodes,
      offlineNodeNumber,
      getNodeMenu,
      handleMousedown,
      nodePosition,
    };
  },
});
</script>

<style scoped>
.node-display {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding-top: 16px;
}

.circle-container {
  position: relative;
  width: 300px;
  height: 300px;
  margin-top: 12px;
}

.circle-node {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 80px;
  height: 80px;
  line-height: 80px;
  text-align: center;
  border-radius: 50%;
  background-color: #1890ff;
  color: #fff;
  font-weight: bold;
  cursor: pointer;
}

.circle-node.offline {
  background-color: #f5222d;
}
</style>
