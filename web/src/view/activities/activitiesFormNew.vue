<template>
  <div>
    <el-row :gutter="15">
      <el-form ref="Union" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-col :span="24">
          <el-form-item label="活动名称" prop="name">
            <el-input v-model="formData.name" placeholder="请输入活动名称" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="活动时间" prop="time">
            <el-date-picker type="daterange" v-model="formData.time" format="yyyy-MM-dd"
                            value-format="yyyy-MM-dd" :style="{width: '100%'}" start-placeholder="开始日期"
                            end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="活动位置" prop="loaction">
            <el-input v-model="formData.loaction" placeholder="请输入活动位置" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="需要人数" prop="neededPersonnel">
            <el-input v-model="formData.neededPersonnel" placeholder="请输入需要人数" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="活动经费" prop="budget">
            <el-input v-model="formData.budget" type="textarea" placeholder="请输入活动经费"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="活动说明" prop="description">
            <el-input v-model="formData.description" type="textarea" placeholder="请输入活动说明"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="申请人" prop="createdBy">
            <el-input v-model="formData.createdBy" placeholder="请输入申请人" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="申请部门" prop="reqUnion">
            <el-select v-model="formData.reqUnion" placeholder="请选择申请部门" clearable :style="{width: '100%'}">
              <!--              <el-option v-for="(item, index) in reqUnionOptions" :key="index" :label="item.label"-->
              <!--                :value="item.value" :disabled="item.disabled"></el-option>-->
              <el-option v-for="(item,key) in unionOptions" :key="key" :label="item.label"
                         :value="item.value"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="审核意见" prop="managementAudit">
            <el-input v-model="formData.managementAudit" placeholder="请输入审核意见" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24" v-auth="888">
          <el-form-item label="申请人ID" prop="createdUserUuid">
            <el-input v-model="formData.createdUserUuid" placeholder="请输入申请人ID" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item size="large">
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-col>
      </el-form>
    </el-row>
  </div>
</template>
<script>
import {createActivities} from "@/api/activities";
import {getUserInfo} from "@/api/user";
import infoList from "@/mixins/infoList";

export default {
  components: {},
  props: [],
  mixins: [infoList],
  data() {

    const validateNumber = (rule, value, callback) => {
      let regex_phone = /^\d+$/g;
      if (!regex_phone.test(value)) {
        return callback(new Error("请输入正整数"));
      }
      return callback();
    };


    return {
      unionOptions: [],
      user_nick_name: "",
      user_uuid: "",
      formData: {
        name: undefined,
        time: null,
        loaction: undefined,
        neededPersonnel: undefined,
        budget: undefined,
        description: undefined,
        createdBy: undefined,
        reqUnion: undefined,
        managementAudit: undefined,
        createdUserUuid: undefined,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入活动名称',
          trigger: 'blur'
        }],
        time: [{
          required: true,
          message: '活动时间不能为空',
          trigger: 'change'
        }],
        loaction: [{
          required: true,
          message: '请输入活动位置',
          trigger: 'blur'
        }],
        neededPersonnel: [{
          required: true,
          message: '请输入需要人数',
          trigger: 'blur'
        }, {
          validator: validateNumber,
          trigger: "blur"
        }],
        budget: [{
          required: true,
          message: '请输入活动经费',
          trigger: 'blur'
        }],
        description: [{
          required: true,
          message: '请输入活动说明',
          trigger: 'blur'
        }],
        createdBy: [{
          required: true,
          message: '请输入申请人',
          trigger: 'blur'
        }],
        reqUnion: [{
          required: true,
          message: '请选择申请部门',
          trigger: 'change'
        }],
        managementAudit: [],
        createdUserUuid: [],
      },
      reqUnionOptions: [{
        "label": "1",
        "value": ""
      }, {
        "label": "2",
        "value": ""
      }, {
        "label": "3",
        "value": ""
      }],
    }
  },
  computed: {

  },
  watch: {},
  async created() {
    await this.getDict("union");

    const userInfo = await getUserInfo();
    const userUUID = userInfo.data[0].uuid;
    const userNickName = userInfo.data[0].nick_name;
    // console.log("UserUUID: " + userUUID);
    // console.log("UserUUIDNickName: " + userNickName);
    this.formData.createdUserUuid = userUUID;
    this.user_uuid = userUUID;
    this.formData.createdBy = userNickName;
    this.user_nick_name = userNickName
  },
  mounted() {
  },
  methods: {
    submitForm() {
      this.$refs['Union'].validate(async valid => {
        if (!valid) return
        this.formData.time = this.formData.time[0] + " - " + this.formData.time[1]
        let res = await createActivities(this.formData);
        if (res.code === 0) {
          this.$message({
            type: "success",
            message: "创建/更改成功"
          })
          this.$router.go(-1)
        }
      })
    },
    resetForm() {
      this.$refs['Union'].resetFields()
    },
  }
}

</script>
<style>
</style>
