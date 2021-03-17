<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="活动名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="活动开始时间">
          <!--          <el-input placeholder="搜索条件" v-model="searchInfo.start_time"></el-input>-->
          <el-date-picker type="date" placeholder="选择开始时间" v-model="searchInfo.start_time" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="活动结束时间">
          <!--          <el-input placeholder="搜索条件" v-model="searchInfo.end_time"></el-input>-->
          <el-date-picker type="date" placeholder="选择结束时间" v-model="searchInfo.end_time" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="审核结果" prop="approved">
          <el-select v-model="searchInfo.approved" clear placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true">
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <!--        <el-form-item>-->
        <!--          <el-button @click="openDialog" type="primary">活动申请</el-button>-->
        <!--        </el-form-item>-->
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
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
      <el-table-column label="发布日期" width="100">
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

      <el-table-column label="活动说明" prop="description" width="120"></el-table-column>

      <el-table-column label="申请人" prop="createdBy" width="120"></el-table-column>

      <el-table-column label="申请部门" prop="reqUnion" width="120">
        <template slot-scope="scope">
          {{ filterDict(scope.row.reqUnion, "union") }}
        </template>
      </el-table-column>

      <el-table-column label="审核结果" prop="approved" width="120" sortable>
        <template slot-scope="scope">{{ checkApproved(scope.row.approved, scope.row.managementAudit) }}</template>
      </el-table-column>

      <el-table-column label="审核意见" prop="managementAudit" width="120"></el-table-column>

<!--      <el-table-column label="申请人ID" prop="createdUserUuid" width="120"></el-table-column>-->

      <el-table-column label="审核" sortable>
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateActivitiesManagement(scope.row)" size="small" type="primary"
                     icon="el-icon-edit">{{ formatAudit(scope.row.approved, scope.row.managementAudit) }}
          </el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="审核" >
      <el-form :model="formData" label-position="right" label-width="80px">
        <!--         <el-form-item label="活动名称:">-->
        <!--            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="活动开始时间:">-->
        <!--              <el-date-picker type="date" placeholder="选择日期" v-model="formData.start_time" clearable></el-date-picker>-->
        <!--       </el-form-item>-->

        <!--         <el-form-item label="活动开始时间:">-->
        <!--              <el-date-picker type="date" placeholder="选择日期" v-model="formData.end_time" clearable></el-date-picker>-->
        <!--       </el-form-item>-->

        <!--         <el-form-item label="活动位置:">-->
        <!--            <el-input v-model="formData.loaction" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="需要人数:"><el-input v-model.number="formData.neededPersonnel" clearable placeholder="请输入"></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="活动经费:">-->
        <!--            <el-input v-model="formData.budget" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="活动说明:">-->
        <!--            <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="申请人:">-->
        <!--            <el-input v-model="formData.createdBy" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->

        <!--         <el-form-item label="申请部门:"><el-input v-model.number="formData.reqUnion" clearable placeholder="请输入"></el-input>-->
        <!--      </el-form-item>-->

        <el-form-item label="审核结果:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                     v-model="formData.approved" clearable></el-switch>
        </el-form-item>

        <el-form-item label="审核意见:">
          <el-input v-model="formData.managementAudit" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <!--         <el-form-item label="申请人ID:">-->
        <!--            <el-input v-model="formData.createdUserUuid" clearable placeholder="请输入" ></el-input>-->
        <!--      </el-form-item>-->
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
  createActivitiesManagement,
  deleteActivitiesManagement,
  deleteActivitiesManagementByIds,
  findActivitiesManagement,
  getActivitiesManagementList,
  updateActivitiesManagement
} from "@/api/ActivitiesManagement"; //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import {getUserInfo} from "@/api/user";

export default {
  name: "ActivitiesManagement",
  mixins: [infoList],
  data() {
    return {
      listApi: getActivitiesManagementList,
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
    formatAudit: (approved, mangAudit) => {
      if (!approved && mangAudit === "") return "审核"
      else return "修改"
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
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.$router.push({name: "ActivityForm"});
      // this.type = "create";
      // this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    await this.getDict("union");

    const userInfo = await getUserInfo();
    this.user_uuid = userInfo.data[0].uuid;

  }
};
</script>

<style>
</style>
