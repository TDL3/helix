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
          <el-form-item label="活动时间" prop="date_picker">
            <el-date-picker @blur="setDate" type="daterange" v-model="date_picker" format="yyyy-MM-dd"
                            :style="{width: '100%'}" start-placeholder="开始日期"
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
            <el-input v-model.number="formData.neededPersonnel" placeholder="请输入需要人数" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <br>
        <el-col :span="8">
          <el-form-item label="活动分数:">
            <el-input v-model.number="formData.score" clearable placeholder="请输入"></el-input>
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
              <!--                         :value="item.value" :disabled="item.disabled"></el-option>-->
              <el-option v-for="(item,key) in unionOptions" :key="key" :label="item.label"
                         :value="item.value"></el-option>
            </el-select>

          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="审核意见" prop="managementAudit" v-auth="888">
            <el-input v-model="formData.managementAudit" placeholder="请输入审核意见" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="审核结果" prop="approved" v-auth="888">
            <el-switch v-model="formData.approved" active-text="批准" inactive-text="否决"></el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="申请人ID" prop="createdUserUuid" v-auth="888">
            <el-input v-model="formData.createdUserUuid" placeholder="请输入申请人ID" :disabled="true"
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

import {getUserInfo} from "@/api/user";
import infoList from "@/mixins/infoList";
import {createActivitiesManagement} from "@/api/ActivitiesManagement";

export default {
  components: {},
  props: [],
  mixins: [infoList],
  data() {
    const validatePersonnel = (rule, value, callback) => {
      if (!/^\d+$/g.test(value)) {
        return callback(new Error("请输入正整数"));
      }
      return callback();
    };
    const validateDate = (rule, value, callback) => {
      // console.log(2)
      if (!this.formData.start_time || !this.formData.end_time) {
        return callback(new Error("请输入日期"));
      }
      return callback();
    };

    return {
      unionOptions: [],
      user_nick_name: "",
      user_uuid: "",
      date_picker: undefined,
      date_picker_validator: undefined,
      formData: {
        name: "",
        start_time: new Date(),
        end_time: new Date(),
        loaction: "",
        neededPersonnel: 0,
        budget: "",
        description: "",
        createdBy: "",
        reqUnion: 0,
        approved: false,
        managementAudit: "",
        createdUserUuid: "",
        score: 0,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入活动名称',
          trigger: 'blur'
        }],
        date_picker: [{
          required: false,
          message: '活动时间不能为空',
          trigger: 'blur'
        }, {
          validator: validateDate,
          trigger: "blur"
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
          validator: validatePersonnel,
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
        approved: [],
        createdUserUuid: [],
      }
    }
  },
  computed: {},
  watch: {},
  async created() {
    await this.getDict("union");

    const userInfo = await getUserInfo();
    const userUUID = userInfo.data[0].uuid;
    const userNickName = userInfo.data[0].nick_name;
    // console.log(userInfo.data[0].authority_id);
    this.formData.createdUserUuid = userUUID;
    this.user_uuid = userUUID;
    this.formData.createdBy = userNickName;
    this.user_nick_name = userNickName
  },
  mounted() {
  },
  methods: {
    setDate() {
      // console.log(1)
      this.formData.start_time = this.date_picker[0];
      this.formData.end_time = this.date_picker[1];
    },
    submitForm() {
      // console.log(this.formData.start_time + " - " + this.formData.end_time);
      this.$refs['Union'].validate(async valid => {
        if (!valid) return
        // this.formData.time = this.formData.time[0] + " - " + this.formData.time[1]
        let res = await createActivitiesManagement(this.formData);
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
      this.$refs['Union'].resetFields();
      this.formData.createdBy = this.user_nick_name;
      this.formData.createdUserUuid = this.user_uuid;
    },
  }
}

</script>
<style>
</style>
