<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="系部" min-width="150" prop="reqUnion">
        <template slot-scope="scope">
          {{ filterDict(scope.row.reqUnion, "union") }}
        </template>
      </el-table-column>
      <el-table-column label="活动得分" min-width="150" prop="score"></el-table-column>
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
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {getUnionScoreList} from "@/api/user";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";

export default {
  mixins: [infoList],

  data() {
    return {
      listApi: getUnionScoreList,
      path: path,
      userInfo: {
        union_id: "",
        score: "",
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {
  },
  async created() {
    await this.getTableData();
    await this.getDict("union")
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
