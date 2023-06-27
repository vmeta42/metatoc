// /* eslint-disable no-async-promise-executor */
// import { create, signup } from "./block";
// import { v4 as uuidv4 } from "uuid";

// const dataInitializationUserList = async () => {
//   return new Promise(async (resolve) => {
//     let userList = localStorage.getItem("userList");
//     let userListArray = JSON.parse(userList);
//     for (const element of userListArray) {
//       if (!element.address && !element.private_key) {
//         let signupResult = await signup();
//         console.log(signupResult);
//         if (
//           signupResult.data &&
//           signupResult.data.address &&
//           signupResult.data.private_key
//         ) {
//           element["address"] = signupResult.data.address;
//           element["private_key"] = signupResult.data.private_key;
//         } else {
//           element["address"] = "";
//           element["private_key"] = "";
//         }
//         localStorage.setItem("userList", JSON.stringify(userListArray));
//       }
//     }
//     localStorage.setItem("currentUser", JSON.stringify(userListArray[0]));
//     resolve();
//   });
// };

// const dataInitializationOnChainChat = () => {
//   return new Promise(async (resolve) => {
//     let userList = localStorage.getItem("userList");
//     let userListArray = JSON.parse(userList);
//     let chat = localStorage.getItem("chatUnconfirm");
//     let chatArray = JSON.parse(chat);
//     let onChainChat = localStorage.getItem("onChainChatUnconfirm");
//     let onChainChatArray = JSON.parse(onChainChat);
//     console.log(
//       "dataInitializationOnChainChat userListArray==>",
//       userListArray
//     );
//     console.log("dataInitializationOnChainChat chatArray==>", chatArray);
//     console.log(
//       "dataInitializationOnChainChat onChainChatArray==>",
//       onChainChatArray
//     );
//     let index = 0;
//     let newChatArray = [];
//     let newOnChainChatArray = [];
//     for (const chatElement of chatArray) {
//       chatElement["uuid"] = uuidv4();
//       if (chatElement["state"] == "on-chain") {
//         // let name = "";
//         let address = "";
//         let private_key = "";
//         let content = {
//           chat: onChainChatArray[index]["chat"],
//           reply: onChainChatArray[index]["reply"],
//         };
//         // let path = "/" + uuidv4();
//         let contentString = JSON.stringify(content);
//         for (const userElement of userListArray) {
//           if (onChainChatArray[index]["users"][0] == userElement["name"]) {
//             // name = userElement["name"];
//             address = userElement["address"];
//             private_key = userElement["private_key"];
//           }
//         }
//         let createResult = await create(
//           private_key,
//           address,
//           "/" + chatElement["uuid"],
//           contentString
//         );
//         if (
//           createResult &&
//           createResult.code == 0 &&
//           createResult.message == "SUCCESSFUL"
//         ) {
//           newChatArray[index] = chatElement;
//           onChainChatArray[index]["uuid"] = chatElement["uuid"];
//           newOnChainChatArray.push(onChainChatArray[index]);
//         } else {
//           chatElement["state"] = "off-chain";
//           newChatArray[index] = chatElement;
//         }
//       } else {
//         newChatArray[index] = chatElement;
//       }
//       localStorage.setItem("chat", JSON.stringify(newChatArray));
//       localStorage.setItem("onChainChat", JSON.stringify(newOnChainChatArray));
//       index++;
//     }
//     localStorage.removeItem("chatUnconfirm");
//     localStorage.removeItem("onChainChatUnconfirm");
//     resolve();
//   });
// };

// const dataInitialization = async () => {
//   // 初始化userList
//   await dataInitializationUserList();
//   // 初始化onChainChat
//   await dataInitializationOnChainChat();

//   localStorage.setItem("initState", "Finish");
// };

// export { dataInitialization };
