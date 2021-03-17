<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="活动名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="活动开始时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.start_time" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="活动开始时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.end_time" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="活动位置:">
                <el-input v-model="formData.loaction" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="需要人数:"><el-input v-model.number="formData.neededPersonnel" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="活动经费:">
                <el-input v-model="formData.budget" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="活动说明:">
                <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="申请人:">
                <el-input v-model="formData.createdBy" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="申请部门:"><el-input v-model.number="formData.reqUnion" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="审核结果:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.approved" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="审核意见:">
                <el-input v-model="formData.managementAudit" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="申请人ID:">
                <el-input v-model="formData.createdUserUuid" clearable placeholder="请输入" ></el-input>
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
    createActivitiesManagement,
    updateActivitiesManagement,
    findActivitiesManagement
} from "@/api/ActivitiesManagement";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "ActivitiesManagement",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            name:"",
            start_time:new Date(),
            end_time:new Date(),
            loaction:"",
            neededPersonnel:0,
            budget:"",
            description:"",
            createdBy:"",
            reqUnion:0,
            approved:false,
            managementAudit:"",
            createdUserUuid:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createActivitiesManagement(this.formData);
          break;
        case "update":
          res = await updateActivitiesManagement(this.formData);
          break;
        default:
          res = await createActivitiesManagement(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findActivitiesManagement({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reacm
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>