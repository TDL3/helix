<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="活动名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="活动时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
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
           
             <el-form-item label="申请部门:">
                 <el-select v-model="formData.reqUnion" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in unionOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
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
    createActivities,
    updateActivities,
    findActivities
} from "@/api/activities";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Activities",
  mixins: [infoList],
  data() {
    return {
      type: "",
      unionOptions:[],
          formData: {
            name:"",
            time:new Date(),
            loaction:"",
            neededPersonnel:0,
            budget:"",
            description:"",
            createdBy:"",
            reqUnion:0,
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
          res = await createActivities(this.formData);
          break;
        case "update":
          res = await updateActivities(this.formData);
          break;
        default:
          res = await createActivities(this.formData);
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
    const res = await findActivities({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reacti
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
    await this.getDict("union");
    
}
};
</script>

<style>
</style>