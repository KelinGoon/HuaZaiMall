<template>
  <div class="app-container">
    <el-card>
      <!--      标题以及按钮-->
      <div slot="header" class="clearfix">
        <el-row
          justify="space-between"
          type="flex"
        >
          <el-col :span="8"><span style="font-size: 24px;">商品管理</span></el-col>
          <el-col :span="2">
            <el-button
              :loading="isLoading"
              icon="el-icon-circle-plus-outline"
              round
              type="success"
              @click="Add()"
            >添加
            </el-button>
          </el-col>
        </el-row>
      </div>
      <!--      标题以及按钮-->

      <!--      添加网站信息的子页面-->
      <el-dialog :visible.sync="dialogAddFormVisible" title="添加网站信息">
        <el-form ref="form" :model="form" :rules="rules">
          <el-form-item
            :label-width="formLabelWidth"
            label="商品名称"
            prop="ProductName"
          >
            <el-input
              v-model="form.ProductName"
              autocomplete="off"
              placeholder="请输入商品名称"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品定价"
            prop="ProductPrice"
          >
            <el-input
              v-model.number="form.ProductPrice"
              autocomplete="off"
              placeholder="请输入商品定价"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品库存"
            prop="ProductStock"
          >
            <el-input
              v-model.number="form.ProductStock"
              autocomplete="off"
              placeholder="请输入商品库存数量"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品描述"
            prop="ProductDescription"
          >
            <el-input
              v-model="form.ProductDescription"
              autocomplete="off"
              placeholder="请对商品进行描述"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品类别"
            prop="CategoryID"
          >
            <el-input
              v-model.number="form.CategoryID"
              autocomplete="off"
              placeholder="请输入商品属于什么类别"
            />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button icon="el-icon-circle-close" @click="handleAddCancel()">取 消</el-button>
          <el-button
            :loading="isLoading"
            icon="el-icon-circle-check"
            type="primary"
            @click="handleAdd()"
          >保 存
          </el-button>
        </div>
      </el-dialog>
      <!--      添加网站信息-->

      <el-table
        v-loading="listLoading"
        :data="list"
        border
        element-loading-text="Loading"
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="产品名">
          <template slot-scope="scope">
            {{ scope.row.ProductName }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="产品描述" width="230">
          <template slot-scope="scope">
            <span>{{ scope.row.ProductDescription }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="价格/元" width="110">
          <template slot-scope="scope">
            <span>{{ scope.row.ProductPrice }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="库存" width="110">
          <template slot-scope="scope">
            <span>{{ scope.row.ProductStock }}</span>
          </template>
        </el-table-column>
        <el-table-column :formatter="formatDate" align="center" label="创建时间" width="170">
          <template slot-scope="scope">
            <span>{{
              new Date(scope.row.CreationDate).toLocaleString('zh', { hour12: false }).replaceAll('/', '-')
            }}</span>
          </template>
        </el-table-column>
        <el-table-column :formatter="formatDate" align="center" label="更新时间" width="170">
          <template slot-scope="scope">
            <span>{{
              new Date(scope.row.UpdateDate).toLocaleString('zh', { hour12: false }).replaceAll('/', '-')
            }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="产品类别" width="110">
          <template slot-scope="scope">
            <span>{{ scope.row.CategoryID }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" width="150">
          <template slot-scope="scope">
            <el-button
              v-show="!scope.row.isEdit"
              circle
              icon="el-icon-edit"
              size="mini"
              type="primary"
              @click="handleEdit(scope.$index, scope.row)"
            />
            <el-button
              v-show="!scope.row.isEdit"
              circle
              icon="el-icon-delete"
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
            />
          </template>
        </el-table-column>
      </el-table>

      <el-dialog :visible.sync="dialogFormVisible" title="编辑网站信息">
        <el-form ref="form" :model="form" :rules="rules">
          <el-form-item
            :label-width="formLabelWidth"
            label="商品名称"
            prop="ProductName"
          >
            <el-input
              v-model="form.ProductName"
              autocomplete="off"
              placeholder="请输入商品名称"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品定价"
            prop="ProductPrice"
          >
            <el-input
              v-model.number="form.ProductPrice"
              autocomplete="off"
              placeholder="请输入商品定价"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品库存"
            prop="ProductStock"
          >
            <el-input
              v-model.number="form.ProductStock"
              autocomplete="off"
              placeholder="请输入商品库存数量"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品描述"
            prop="ProductDescription"
          >
            <el-input
              v-model="form.ProductDescription"
              autocomplete="off"
              placeholder="请对商品进行描述"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="商品类别"
            prop="CategoryID"
          >
            <el-input
              v-model.number="form.CategoryID"
              autocomplete="off"
              placeholder="请输入商品属于什么类别"
            />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button icon="el-icon-circle-close" @click="handleCancel()">取 消</el-button>
          <el-button
            :loading="isLoading"
            icon="el-icon-circle-check"
            type="primary"
            @click="handleSave()"
          >保 存
          </el-button>
        </div>
      </el-dialog>

      <el-pagination
        :current-page="pageNum"
        :page-size="pageSize"
        :page-sizes="[5, 10, 15, 20]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </el-card>
  </div>
</template>

<script>

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      // siteList: [],
      dialogFormVisible: false,
      dialogAddFormVisible: false,
      isLoading: false,
      formLabelWidth: '120px',
      // 分页参数
      pageSize: 10,
      pageNum: 1,
      total: 0,
      form: {
        // 产品的form
        ProductName: '',
        ProductPrice: '',
        ProductStock: '',
        ProductDescription: '',
        UpdateDate: new Date(),
        CategoryID: 0
      },
      // 表单数据校验规则
      rules: {
        ProductName: [
          { required: true, message: '请输入商品名称', trigger: 'blur' },
          { min: 3, message: '长度不得小于 3 个字符', trigger: 'blur' }
        ],
        ProductPrice: [
          { required: true, message: '请输入商品库存量', trigger: 'blur' }
        ],
        ProductStock: [
          { required: true, message: '请输入商品库存量', trigger: 'blur' }
        ],
        ProductDescription: [
          { required: true, message: '请输入商品描述', trigger: 'blur' }
        ],
        CategoryID: [
          { required: true, message: '请输入商品类别', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getProduct()
  },
  methods: {
    // // get函数
    // getProduct() {
    //   this.listLoading = true
    //   getList().then(response => {
    //     this.list = response.data
    //     this.listLoading = false
    //   })
    // },
    getProduct() {
      this.$store
        .dispatch('product/getList', {
          pagesize: this.pageSize,
          pagenum: this.pageNum
        })
        .then((res) => {
          if (res.status_code === 0) {
            const { data } = res
            this.list = data
            this.total = res.total
            this.listLoading = false
          } else {
            this.$message.error(res.status_msg)
          }
        })
    },
    // 分页控件函数--start
    handleSizeChange(val) {
      this.pageSize = val
      this.getProduct()
    },
    handleCurrentChange(val) {
      this.pageNum = val
      this.getProduct()
    },
    // 分页控件函数--end
    Add() {
      this.dialogAddFormVisible = true
    },
    handleAdd() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          this.isLoading = true
          this.$store.dispatch('product/addList', this.form).then((res) => {
            console.log(res)
            if (res.status_code === 0) {
              this.dialogAddFormVisible = false
              this.$message.success(res.status_msg)
              console.log('dfsdfdsfdsfd')
              this.form = {}
              this.getProduct()
            } else {
              this.$message.error(res.status_msg)
            }
          })
          this.isLoading = false
        } else {
          return false
        }
      })
    },
    handleAddCancel() {
      this.dialogAddFormVisible = false
      this.form = {}
      this.$message.info('取消添加')
    },
    handleEdit(index, row) {
      this.form = row
      this.dialogFormVisible = true
    },
    handleCancel() {
      this.dialogFormVisible = false
      this.getProduct()
      this.form = {}
      this.$message.info('取消编辑')
    },
    handleSave(index, row) {
      this.$refs['form'].validate((valid) => {
        this.form.UpdateDate = new Date() // 使用浏览器默认的本地化格式
        if (valid) {
          this.isLoading = true
          this.$store.dispatch('product/UpdateList', this.form).then((res) => {
            if (res.status_code === 0) {
              this.$message.success(res.status_msg)
              this.dialogFormVisible = false
              this.form = {}
            } else {
              this.$message.error(res.status_msg)
            }
            this.isLoading = false
          })
        } else {
          return false
        }
      })
    },
    handleDelete(index, row) {
      this.$confirm(
        '此操作将删除商品：' +
        row.ProductName +
        ' 是否继续?',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )
        .then(() => {
          this.$store.dispatch('product/DeleteList', row).then((res) => {
            if (res.status_code === 0) {
              this.$message.success(res.status_msg)
              this.dialogFormVisible = false
              this.getProduct()
            } else {
              this.$message.error(res.status_msg)
            }
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    }

  }
}

function time(time = +new Date()) {
  var date = new Date(time + 8 * 3600 * 1000)
  return date.toJSON().substr(0, 19).replace('T', ' ').replace(/-/g, '.')
}

time() // "2018.08.09 18:25:54"
</script>
