import { $post, $put } from "./request";

const create = (private_key, address, path, content) => {
  return $post("/block/paths", {
    private_key: private_key,
    address: address,
    path: path,
    content: content,
  }).then(
    (res) => {
      console.log("post paths resp res==>", JSON.parse(res.data));
      return JSON.parse(res.data);
    },
    (err) => {
      console.log("post paths resp err==>", err);
      return null;
    }
  );
};

const detail = () => {};

const list = () => {};

const share = (private_key, from_address, to_address, token_name) => {
  return $put("/block/paths", {
    private_key,
    from_address,
    to_address,
    token_name,
  }).then(
    (res) => {
      console.log("put share resp res==>", JSON.parse(res.data));
      return JSON.parse(res.data);
    },
    (err) => {
      console.log("put share resp err==>", err);
      return null;
    }
  );
};

const signup = () => {
  return $post("/metatoc-service/v1/blockchain/signup").then(
    (res) => {
      // console.log("signup resp res==>", JSON.parse(res.data));
      return res.data.data;
    },
    (err) => {
      console.log("signup resp err==>", err);
      return null;
    }
  );
};

export { create, detail, list, share, signup };
