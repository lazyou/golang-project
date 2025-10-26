<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联分类:" prop="stockCategoryId">
          <el-select v-model="formData.stockCategoryId" placeholder="请选择关联分类" filterable style="width:100%"
                     :clearable="true">
            <el-option v-for="(item,key) in dataSource.stockCategoryId" :key="key" :label="item.label"
                       :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序(越小越靠前):" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序(越小越靠前)" />
        </el-form-item>
        <el-form-item label="股票名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入股票名字" />
        </el-form-item>
        <el-form-item label="股票代码:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入股票代码" />
        </el-form-item>
        <el-form-item label="公司名称:" prop="companyName">
          <el-input v-model="formData.companyName" :clearable="true" placeholder="请输入公司名称" />
        </el-form-item>
        <el-form-item label="公司简介:" prop="companyProfile">
          <el-input v-model="formData.companyProfile" :clearable="true" placeholder="请输入公司简介" />
        </el-form-item>
        <el-form-item label="所属行业:" prop="industry">
          <el-input v-model="formData.industry" :clearable="true" placeholder="请输入所属行业" />
        </el-form-item>
        <el-form-item label="上市日期:" prop="listingDate">
          <el-date-picker v-model="formData.listingDate" type="date" style="width:100%" placeholder="选择日期"
                          :clearable="true" />
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <RichEdit v-model="formData.remark" />
        </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import {
    getStockItemDataSource,
    createStockItem,
    updateStockItem,
    findStockItem
  } from '@/api/stock/stockItem'

  defineOptions({
    name: 'StockItemForm'
  })

  // 自动获取字典
  import { getDictFunc } from '@/utils/format'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage } from 'element-plus'
  import { ref, reactive } from 'vue'
  // 富文本组件
  import RichEdit from '@/components/richtext/rich-edit.vue'


  const route = useRoute()
  const router = useRouter()

  // 提交按钮loading
  const btnLoading = ref(false)

  const type = ref('')
  const formData = ref({
    stockCategoryId: undefined,
    sort: undefined,
    name: '',
    code: '',
    companyName: '',
    companyProfile: '',
    industry: '',
    listingDate: new Date(),
    remark: ''
  })
  // 验证规则
  const rule = reactive({
    stockCategoryId: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }],
    name: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }],
    code: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }]
  })

  const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async () => {
    const res = await getStockItemDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

  // 初始化方法
  const init = async () => {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findStockItem({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
  }

  init()
  // 保存按钮
  const save = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createStockItem(formData.value)
          break
        case 'update':
          res = await updateStockItem(formData.value)
          break
        default:
          res = await createStockItem(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    })
  }

  // 返回按钮
  const back = () => {
    router.go(-1)
  }

</script>

<style>
</style>
