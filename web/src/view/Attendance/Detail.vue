<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <CustomPic :picSrc="scope.row.headerImg" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="签到情况" min-width="150" prop="present">
        <template slot-scope="scope">{{scope.row.present|formatBoolean}}</template>
      </el-table-column>
      <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{float:'right',padding:'20px'}"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
          layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
    </el-table>
  </div>
</template>


<script>
import {getAttendantList} from "@/api/ActivitiesManagement";

const path = process.env.VUE_APP_BASE_API;
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
import CustomPic from "@/components/customPic";
export default {
  name: "Api",
  mixins: [infoList],
  components: { CustomPic },
  data() {
    return {
      listApi: getAttendantList,
      path: path,
      userInfo: {
        username: "",
        score: "",
        nickName: "",
        headerImg: "",
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {

  },
  filters: {
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "已签到" : "未签到";
      } else {
        return "";
      }
    }
  },
  async created() {
    await this.getTableData({ID: this.$route.query.id});
    // console.log(this.$route.query.id);
    if(this.$route.query.id) {
      const res = await getAttendantList({ID: this.$route.query.id});
      console.log(res);
      let payload = res.data.list;
      if (payload != null) {
        let cache = [];
        payload.forEach(el => {
          el.user.present = el.present
          console.log(el.user)
          cache.push(el.user)
        });
        this.tableData = cache;
      }
    }
  }
};
</script>
<style lang="scss">

.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>
