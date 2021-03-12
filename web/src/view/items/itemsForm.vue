<template>
  <div>
    <el-form :model="formData" :rules="rules" label-position="right" label-width="80px">
<!--      <el-form-item label="标题:">-->
<!--        <el-input v-model="formData.title" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="丢失地点:">-->
<!--        <el-input v-model="formData.location" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="丢失时间:">-->
<!--        <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="已找到:">-->
<!--        <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"-->
<!--                   v-model="formData.isFond" clearable></el-switch>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="失物图片:">-->
<!--        <el-input v-model="formData.picture" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="QQ:">-->
<!--        <el-input v-model="formData.QQ" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="微信:">-->
<!--        <el-input v-model="formData.wechat" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="手机:">-->
<!--        <el-input v-model="formData.phone" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="详细描述:">-->
<!--        <el-input v-model="formData.description" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="创建者:">-->
<!--        <el-input v-model="formData.createdBy" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="uuid:">-->
<!--        <el-input v-model="formData.uuid" clearable placeholder="请输入"></el-input>-->
<!--      </el-form-item>-->
      <el-form-item label="标题:">
        <el-input :rule="rules.title" v-model="formData.title" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="丢失地点:">
        <el-input v-model="formData.location" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="丢失时间:">
        <el-date-picker type="date" value-format="yyyy-MM-dd" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
      </el-form-item>

      <el-form-item label="已找到:">
        <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                   v-model="formData.isFond" clearable></el-switch>
      </el-form-item>

      <el-form-item label="失物图片:">
        <el-input v-model="formData.picture" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="QQ:">
        <el-input v-model="formData.QQ" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="微信:">
        <el-input v-model="formData.wechat" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="手机:">
        <el-input v-model="formData.phone" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="详细描述:">
        <el-input v-model="formData.description" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="创建者:">
        <el-input v-model="formData.createdBy" clearable placeholder="请输入"></el-input>
      </el-form-item>

      <el-form-item label="uuid:">
        <el-input v-model="formData.uuid" clearable placeholder="请输入"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="save" type="primary">保存</el-button>
        <el-button @click="back" type="primary">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
  createItems,
  updateItems,
  findItems, getUserInfo
} from "@/api/items";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";

export default {
  name: "Items",
  mixins: [infoList],
  data() {
    return {
      type: "",
      formData: {
        title: undefined,
        location: "",
        time: "",
        isFond: false,
        picture: "",
        QQ: "",
        wechat: "",
        phone: "",
        description: "",
        createdBy: "",
        uuid: "",

      },
      rules: {
        title: [
          { required: true, message: "请输入标题", trigger: "blur" },
        ],
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createItems(this.formData);
          break;
        case "update":
          res = await updateItems(this.formData);
          break;
        default:
          res = await createItems(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findItems({ID: this.$route.query.id})
      if (res.code == 0) {
        this.formData = res.data.reitm
        this.type = "update"
      }
    } else {
      this.type = "create"
    }
    const userInfo = await getUserInfo();
    const userUUID = userInfo.data[0].uuid;
    const userNickName = userInfo.data[0].nick_name;
    console.log("UserUUID: " + userUUID);
    console.log("UserUUIDNickName: " + userNickName);
    this.formData.uuid = userUUID;
    this.user_uuid = userUUID;
    this.formData.createdBy = userNickName;
    this.user_nick_name = userNickName
  }
};
</script>

<style>
</style>
