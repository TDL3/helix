<template>
  <div>
    <el-form ref="NewItemForm" :model="formData" :rules="rules" size="medium" label-width="100px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="formData.title" placeholder="请输入标题" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="丢失位置" prop="location">
        <el-input v-model="formData.location" placeholder="请输入丢失位置" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="丢失时间" prop="time">
        <el-date-picker v-model="formData.time" format="yyyy-MM-dd" value-format="yyyy-MM-dd"
                        :style="{width: '100%'}" placeholder="请选择丢失时间" clearable></el-date-picker>
      </el-form-item>
      <el-form-item label="已找到" prop="isFond">
        <el-switch v-model="formData.isFond"></el-switch>
      </el-form-item>

      <el-form-item label="失物图片" prop="picture">
        <upload-image v-model="formData.picture" :fileSize="512" :maxWH="1080"/>
      </el-form-item>

      <el-form-item label="QQ" prop="qq">
        <el-input v-model="formData.qq" placeholder="请输入QQ" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="微信" prop="wechat">
        <el-input v-model="formData.wechat" placeholder="请输入微信" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="电话" prop="phone">
        <el-input v-model="formData.phone" placeholder="请输入电话" clearable :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="详细描述" prop="description">
        <el-input v-model="formData.description" type="textarea" placeholder="请输入详细描述"
                  :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
      </el-form-item>
<!--      <el-form-item label="创建者" prop="createdBy">-->
<!--        <el-input v-model="formData.createdBy" placeholder="请输入创建者" clearable :style="{width: '100%'}">-->
<!--        </el-input>-->
<!--      </el-form-item>-->
<!--      <el-form-item label="UUID" prop="uuid">-->
<!--        <el-input v-model="formData.uuid" placeholder="请输入UUID" clearable :style="{width: '100%'}"></el-input>-->
<!--      </el-form-item>-->
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>

import UploadImage from "@/components/upload/image";
import {createItems, getUserInfo} from "@/api/items";

const path = process.env.VUE_APP_BASE_API;
export default {
  components: {
    UploadImage
  },
  props: [],
  data() {
    const validatePhone = (rule, value, callback) => {
      let regex_phone = /(\+?( |-|\.)?\d{1,2}( |-|\.)?)?(\(?\d{3}\)?|\d{3})( |-|\.)?(\d{3}( |-|\.)?\d{4})/g;
      if (!regex_phone.test(value)) {
        return callback(new Error("请输入正确的电话"));
      }
      return callback();
    };
    const validateQQ = (rule, value, callback) => {
      if (!/[1-9][0-9]{4,}/g.test(value)) {
        return callback(new Error("请输入正确的QQ"));
      }
      return callback();
    };
    return {
      path: path,
      imageUrl: "",
      formData: {
        title: undefined,
        location: undefined,
        time: null,
        isFond: false,
        picture: null,
        qq: undefined,
        wechat: undefined,
        phone: undefined,
        description: undefined,
        createdBy: undefined,
        uuid: undefined,
        approved:false,
      },
      rules: {
        title: [{
          required: true,
          message: '请输入标题',
          trigger: 'blur'
        }],
        location: [{
          required: true,
          message: '请输入丢失位置',
          trigger: 'blur'
        }],
        time: [{
          required: true,
          message: '请选择丢失时间',
          trigger: 'change'
        }],
        qq: [{
          validator: validateQQ,
          trigger: "blur"
        }],
        wechat: [],
        phone: [{
          required: true,
          message: '请输入电话',
          trigger: 'blur'
        }, {
          validator: validatePhone,
          trigger: "blur"
        }],
        description: [{
          required: true,
          message: '请输入详细描述',
          trigger: 'blur'
        }],
        createdBy: [{
          required: true,
          message: '请输入创建者',
          trigger: 'blur'
        }],
        uuid: [{
          required: true,
          message: '请输入UUID',
          trigger: 'blur'
        }],
      },
      pictureAction: `${path}/fileUploadAndDownload/upload`,
      picturefileList: [],
    }
  },
  computed: {},
  watch: {},
  async created() {
    const userInfo = await getUserInfo();
    const userUUID = userInfo.data[0].uuid;
    const userNickName = userInfo.data[0].nick_name;
    // console.log("UserUUID: " + userUUID);
    // console.log("UserUUIDNickName: " + userNickName);
    this.formData.uuid = userUUID;
    this.user_uuid = userUUID;
    this.formData.createdBy = userNickName;
    this.user_nick_name = userNickName
  },
  mounted() {
  },
  methods: {
    submitForm() {
      this.$refs['NewItemForm'].validate(async (valid) => {
        if (!valid) return
        let res = await createItems(this.formData);
        if (res.code === 0) {
          this.$message({
            type: "success",
            message: "创建/更改成功"
          })
        }
      });
      this.$router.go(-1)
    },
    resetForm() {
      this.$refs['NewItemForm'].resetFields()
    },


    checkFile(file) {
      this.fullscreenLoading = true;
      const isJPG = file.type === "image/jpeg";
      const isPng = file.type === "image/png";
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isJPG && !isPng) {
        this.$message.error("上传头像图片只能是 JPG或png 格式!");
        this.fullscreenLoading = false;
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
        this.fullscreenLoading = false;
      }
      return (isPng || isJPG) && isLt2M;
    },
    uploadSuccess(res) {
      this.fullscreenLoading = false;
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "上传成功"
        });
        if (res.code == 0) {
          this.getTableData();
        }
      } else {
        this.$message({
          type: "warning",
          message: res.msg
        });
      }
    },
    uploadError() {
      this.$message({
        type: "error",
        message: "上传失败"
      });
      this.fullscreenLoading = false;
    }
  }
}

</script>
<style>
.el-upload__tip {
  line-height: 1.2;
}

</style>
