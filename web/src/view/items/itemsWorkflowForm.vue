<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="标题:">
                <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="丢失地点:">
                <el-input v-model="formData.location" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="丢失时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
           </el-form-item>

             <el-form-item label="已找到:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isFond" clearable ></el-switch>
          </el-form-item>

             <el-form-item label="失物图片:">
                <el-input v-model="formData.picture" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="QQ:">
                <el-input v-model="formData.QQ" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="微信:">
                <el-input v-model="formData.wechat" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="手机:">
                <el-input v-model="formData.phone" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="详细描述:">
                <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="创建者:">
                <el-input v-model="formData.createdBy" clearable placeholder="请输入" ></el-input>
          </el-form-item>

             <el-form-item label="uuid:">
                <el-input v-model="formData.uuid" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button v-if="this.wf.clazz == 'start'" @click="start" type="primary">启动</el-button>
           <!-- complete传入流转参数 决定下一步会流转到什么位置 此处可以设置多个按钮来做不同的流转 -->
           <el-button v-if="canShow" @click="complete('yes')" type="primary">提交</el-button>
           <el-button v-if="showSelfNode" @click="complete('')" type="primary">确认</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    startWorkflow,
    completeWorkflowMove
} from "@/api/workflowProcess";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Items",
  mixins: [infoList],
  props:{
      business:{
         type:Object,
        default:function(){return null}
      },
      wf:{
        type:Object,
        default:function(){return{}}
      },
      move:{
         type:Object,
         default:function(){return{}}
      },
      workflowMoveID:{
        type:[Number,String],
        default:0
      }
   },
  data() {
    return {formData: {
            title:"",
            location:"",
            time:new Date(),
            isFond:false,
            picture:"",
            QQ:"",
            wechat:"",
            phone:"",
            description:"",
            createdBy:"",
            uuid:"",

      }
    };
  },
  computed:{
      showSelfNode(){
         if(this.wf.assignType == "self" && this.move.promoterID == this.userInfo.ID){
             return true
         }else{
             return false
         }
      },
      canShow(){
         if(this.wf.assignType == "user"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.ID+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }else if(this.wf.assign_type == "authority"){
            if(this.wf.assignValue.indexOf(","+this.userInfo.authorityId+",")>-1 && this.wf.clazz == 'userTask'){
               return true
            }else{
               return false
            }
         }
         return false
      },
      ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessId:0,
              businessType:"itm",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"create",
              param:""
              }
          });
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"启动成功"
        })
       this.back()
      }
    },
    async complete(param){
     const res = await completeWorkflowMove({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessID:this.formData.ID,
              businessType:"itm",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"complete",
              param:param
              }
     })
     if(res.code == 0){
       this.$message({
          type:"success",
          message:"提交成功"
       })
       this.back()
     }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
    if(this.business){
     this.formData = this.business
    }
}
};
</script>

<style>
</style>
