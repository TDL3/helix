<template>
  <div>
    <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="发布日期" width="110" sortable>
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>

      <el-table-column label="活动名称" prop="name" width="120"></el-table-column>

      <!--    <el-table-column label="活动开始时间" prop="start_time" width="120">-->
      <!--      <template slot-scope="scope">{{scope.row.start_time|formatDate}}</template>-->
      <!--    </el-table-column>-->

      <!--    <el-table-column label="活动结束时间" prop="end_time" width="120">-->
      <!--      <template slot-scope="scope">{{scope.row.end_time|formatDate}}</template>-->
      <!--    </el-table-column>-->
      <el-table-column label="活动时间" prop="end_time" width="120">
        <template slot-scope="scope">{{ scope.row.start_time|formatDate }} 至 {{ scope.row.end_time|formatDate }}
        </template>
      </el-table-column>

      <el-table-column label="活动位置" prop="loaction" width="120"></el-table-column>

      <el-table-column label="需要人数" prop="neededPersonnel" width="120"></el-table-column>

      <el-table-column label="活动经费" prop="budget" width="120"></el-table-column>

      <el-table-column label="活动说明" prop="description" width="380"></el-table-column>


      <el-table-column label="状态">
        <template slot-scope="scope">
          <span v-if="user_uuid===scope.row.createdUserUuid">
            <el-button class="table-button" @click="updateActivitiesManagement(scope.row)" size="small" type="primary"
                       icon="el-icon-edit">参加</el-button>
          </span>
        </template>
      </el-table-column>
    </el-table>

<!--    <el-pagination-->
<!--        :current-page="page"-->
<!--        :page-size="pageSize"-->
<!--        :page-sizes="[10, 30, 50, 100]"-->
<!--        :style="{float:'right',padding:'20px'}"-->
<!--        :total="total"-->
<!--        @current-change="handleCurrentChange"-->
<!--        @size-change="handleSizeChange"-->
<!--        layout="total, sizes, prev, pager, next, jumper"-->
<!--    ></el-pagination>-->

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="确定加入吗？">
      <el-form :model="formData" label-position="right" label-width="20px">

      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  deleteActivitiesManagement,
  deleteActivitiesManagementByIds,
  findActivitiesManagement,
  getAttendedActivitiesInfoList,
} from "@/api/ActivitiesManagement"; //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {getUserInfo} from "@/api/user";

export default {
  name: "ActivitiesManagement",
  mixins: [infoList],
  data() {
    return {
      listApi: getAttendedActivitiesInfoList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      unionOptions: [],
      user_uuid: "",
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

      }
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    }
  },
  methods: {
    checkApproved: (approved, mangAudit) => {
      if (approved) return "批准"
      if (!approved && mangAudit === "") return "审核中"
      if (!approved && mangAudit !== "") return "否决"
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.approved == "") {
        this.searchInfo.approved = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteActivitiesManagement(row);
      });
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length == 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteActivitiesManagementByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length == ids.length) {
          this.page--;
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateActivitiesManagement(row) {
      const res = await findActivitiesManagement({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reacm;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
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

      };
    },
    async deleteActivitiesManagement(row) {
      const res = await deleteActivitiesManagement({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
          this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      // switch (this.type) {
      //   case "create":
      //     res = await createActivitiesManagement(this.formData);
      //     break;
      //   case "update":
      //     res = await updateActivitiesManagement(this.formData);
      //     break;
      //   default:
      //     res = await createActivitiesManagement(this.formData);
      //     break;
      // }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        await this.getTableData();
      }
    },
    openDialog() {
      this.$router.push({name: "ActivityForm"});
      // this.type = "create";
      // this.dialogFormVisible = true;
    }
  },
  async created() {
    this.searchInfo.approved = 1;
    await this.getTableData();
    await this.getDict("union");

    const userInfo = await getUserInfo();
    this.user_uuid = userInfo.data[0].uuid;

  }
};
</script>

<style>
</style>
