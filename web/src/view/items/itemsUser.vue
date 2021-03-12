<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="标题">
          <el-input placeholder="搜索条件" v-model="searchInfo.title"></el-input>
        </el-form-item>
        <el-form-item label="丢失地点">
          <el-input placeholder="搜索条件" v-model="searchInfo.location"></el-input>
        </el-form-item>
        <el-form-item label="丢失时间">
          <el-date-picker type="date" value-format="yyyy-MM-dd" placeholder="选择日期" v-model="searchInfo.time" clearable></el-date-picker>
          <!--          <el-input placeholder="搜索条件" v-model="searchInfo.time"></el-input>-->
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">添加失物招领信息</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
<!--            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>-->
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
      <el-table-column label="创建日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>

      <el-table-column label="标题" prop="title" width="120"></el-table-column>

      <el-table-column label="丢失地点" prop="location" width="120"></el-table-column>

      <el-table-column label="丢失时间" prop="time" width="120"></el-table-column>

      <el-table-column label="已找到" prop="isFond" width="120">
        <template slot-scope="scope">{{ scope.row.isFond|formatBoolean }}</template>
      </el-table-column>

      <el-table-column label="失物图片" prop="picture" width="240" >
        <template slot-scope="scope">
          <img class="image" @click="openImgInNewTab(path + scope.row.picture)" :src="path + scope.row.picture" alt="失物图片">
        </template>
      </el-table-column>

      <el-table-column label="QQ" prop="QQ" width="120"></el-table-column>

      <el-table-column label="微信" prop="wechat" width="120"></el-table-column>

      <el-table-column label="手机" prop="phone" width="120"></el-table-column>

      <el-table-column label="详细描述" prop="description" width="200"></el-table-column>

      <el-table-column label="创建者" prop="createdBy" width="120"></el-table-column>

<!--      <el-table-column label="uuid" prop="uuid" width="120"></el-table-column>-->

<!--      <el-table-column label="操作">-->
<!--        <template slot-scope="scope">-->
<!--          <el-button class="table-button" @click="updateItems(scope.row)" size="small" type="primary"-->
<!--                     icon="el-icon-edit">变更-->
<!--          </el-button>-->
<!--          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>-->
<!--        </template>-->
<!--      </el-table-column>-->
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="物品信息">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="丢失地点:">
          <el-input v-model="formData.location" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="丢失时间:">
          <el-date-picker type="date" value-format="yyyy-MM-dd" placeholder="选择日期" v-model="formData.time" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="已找到:">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                     v-model="formData.isFond" clearable></el-switch>
        </el-form-item>

        <el-form-item label="失物图片" prop="picture">
          <upload-image v-model="formData.picture" :fileSize="512" :maxWH="1080" />
        </el-form-item>

        <el-form-item label="QQ:">
          <el-input v-model="formData.QQ" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="微信:">
          <el-input v-model="formData.wechat" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="手机:">
          <el-input v-model="formData.phone" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="详细描述:">
          <el-input v-model="formData.description" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="创建者:">
          <el-input v-model="formData.createdBy" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="uuid:">
          <el-input v-model="formData.uuid" clearable placeholder="请输入"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
// console.log(path)
import {
  createItems,
  deleteItems,
  deleteItemsByIds,
  updateItems,
  findItems,
  getItemsList,
  getUserInfo
} from "@/api/items";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";
import UploadImage from "@/components/upload/image";

export default {

  name: "Items",
  mixins: [infoList],
  data() {
    return {
      path: path,
      listApi: getItemsList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      user_uuid: "",
      user_nick_name: "",
      multipleSelection: [], formData: {
        title: "",
        location: "",
        time: "",
        isFond: false,
        picture: "",
        QQ: "",
        wechat: "",
        phone: "",
        description: "",
        createdBy: "",
        uuid: "",

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
  components: {
    UploadImage
  },
  methods: {
    openImgInNewTab(url) {
      let windowObjectReference = window.open(
          url,
          "大图"
      );
      windowObjectReference.focus();
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.isFond == "") {
        this.searchInfo.isFond = null
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
        this.deleteItems(row);
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
      const res = await deleteItemsByIds({ids})
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
    async updateItems(row) {
      const res = await findItems({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reitm;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        title: "",
        location: "",
        time: "",
        isFond: false,
        picture: "",
        QQ: "",
        wechat: "",
        phone: "",
        description: "",
        createdBy: this.user_nick_name,
        uuid: this.user_uuid,

      };
    },
    async deleteItems(row) {
      const res = await deleteItems({ID: row.ID});
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
          res = await createItems(this.formData);
          break;
        case "update":
          res = await updateItems(this.formData);
          break;
        default:
          res = await createItems(this.formData);
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
      this.type = "create";
      this.$router.push({name: "NewItemForm"});
      // this.dialogFormVisible = false;
    }
  },
  async created() {
    await this.getTableData();
    const userInfo = await getUserInfo();
    const userUUID = userInfo.data[0].uuid;
    const userNickName = userInfo.data[0].nick_name;
    // console.log("UserUUID: " + userUUID);
    // console.log("UserUUIDNickName: " + userNickName);
    this.formData.uuid = userUUID;
    this.user_uuid = userUUID;
    this.formData.createdBy = userNickName;
    this.user_nick_name = userNickName

  }
};
</script>

<style scoped>
.image {
  /*width: 158px;*/
  height: 118px;
  /*display: block;*/
}
.image:hover {
  cursor: pointer;
}
</style>
