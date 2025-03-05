<template>
  <div class="app-container">
    <div class="search-container">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item prop="value" label="配置值">
          <el-input v-model="queryParams.value" placeholder="配置值" clearable @keyup.enter="fetchPage" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="fetchPage"><i-ep-search />搜索</el-button>
          <el-button @click="resetQuery"><i-ep-refresh />重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never" class="table-container">
        <template #header>
            <el-button @click="openDialogWithAdd()" type="success"><i-ep-plus />新增</el-button>
            <el-button type="danger" :disabled="state.ids.length === 0" @click="bantchDelete()"><i-ep-delete />删除</el-button>
        </template>
      <el-table ref="dataTableRef" v-loading="state.loading" :data="configItems" highlight-current-row border
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="ID" prop="id" align="center" />
        <el-table-column label="配置名称" prop="name" align="center" />
<el-table-column label="配置值" prop="value" align="center" />
        <el-table-column fixed="right" label="操作" width="140" align="center">
          <template #default="{row}">
            <el-button type="primary" size="small" link @click="openDialogWithEdit(row.id)">
              <i-ep-edit />编辑
            </el-button>
            <el-button type="primary" size="small" link @click="handleDelete(row.id)">
              <i-ep-delete />删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination v-if="state.total > 0" v-model:total="state.total" v-model:page="queryParams.page"
        v-model:limit="queryParams.size" @pagination="fetchPage" />
    </el-card>

    <!-- Config 表单弹窗 -->
    <el-dialog v-model="state.dialogVisible" :title="state.dialogTitle" @close="closeDialog">
      <el-form ref="configFormRef" :model="formData" :rules="rules" label-width="100px">
        <el-form-item label="ID" prop="id" v-if="formData.id>0">
          <el-input v-model="formData.id" disabled placeholder="" />
        </el-form-item>
        <el-form-item prop="name" label="配置名称">
          <el-input v-model="formData.name" placeholder="配置名称" clearable />
        </el-form-item>
<el-form-item prop="value" label="配置值">
          <el-input v-model="formData.value" placeholder="配置值" clearable />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="handleSubmit">确 定</el-button>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">

import { setConfigAdd } from "@/api/demo/api";
import { setConfigUpdate } from "@/api/demo/api";
import { deleteConfigDelete } from "@/api/demo/api";
import { getConfigForm } from "@/api/demo/api";
import { getConfigPage } from "@/api/demo/api";

defineOptions({
  name: "config",
  inheritAttrs: false,
});

const queryFormRef = ref(ElForm);

const configFormRef = ref(ElForm);
const configItems = ref<Demo.ConfigItem[]>();
const state = reactive({
  loading: false,
  total: 0,
  ids: [] as number[],
  dialogVisible: false,
  dialogTitle: "",
});

const queryParams = reactive<Demo.ConfigPageReq>({
  page: 1,
  size: 10,
  value: 0,
});

const formData = reactive<Demo.ConfigItem>({
  id: 0,
name: "",
value: 0,
  create_at: "",
  update_at: ""
});

// 根据需要添加校验规则
const rules = reactive({
//   name: [{ required: true, message: "本字段必填", trigger: "blur" }],
});

/** 查询 */
const fetchPage = async () => {
  state.loading = true;
  const rsp = await getConfigPage(queryParams);
  state.loading = false;
  if (rsp.result == 0) {
    configItems.value = rsp.data.items;
    state.total = rsp.data.total;
  }
}
/** 重置查询 */
function resetQuery() {
  queryFormRef.value.resetFields();
  queryParams.page = 1;
  fetchPage();
}

/** 行checkbox 选中事件 */
function handleSelectionChange(selection: any) {
  state.ids = selection.map((item: any) => item.id);
}

/** 打开添加弹窗 */
function openDialogWithAdd() {
  state.dialogVisible = true;
  state.dialogTitle = "添加Config";
  resetForm();
}
/** 打开编辑弹窗 */
const openDialogWithEdit  = async (roleId: number) => {
  state.dialogVisible = true;
  state.dialogTitle = "修改Config";
  state.loading = true;
  const rsp = await getConfigForm({ id: roleId });
  state.loading = false;
  if (rsp.result == 0) {
    Object.assign(formData, rsp.data);
  }
}

/** 保存提交 */
function handleSubmit() {
  configFormRef.value.validate((valid: any) => {
    if (valid) {
      if (formData.id) {
        updateRowRecord();
      } else {
        addRowRecord();
      }
    }
  });
}
  /** 新增记录 */
const addRowRecord = async () => {
    state.loading = true;
    const rsp = await setConfigAdd(formData);
    state.loading = false
    if (rsp.result == 0) {
        ElMessage.success("添加成功");
        closeDialog();
        resetQuery();
    }
}
/** 修改记录 */
const updateRowRecord = async () => {
    state.loading = true;
    const rsp = await setConfigUpdate(formData);
    state.loading = false
    if (rsp.result == 0) {
        ElMessage.success("修改成功");
        closeDialog();
        resetQuery();
    }
}
/** 关闭表单弹窗 */
function closeDialog() {
  state.dialogVisible = false;
  resetForm();
}

/** 重置表单 */
function resetForm() {
    const value = configFormRef.value;
    if(value) {
        value.resetFields();
        value.clearValidate();
    }
    formData.id = 0;
    formData.name = "";
formData.value = 0;
    formData.create_at = "";
    formData.update_at = "";
}

/** 删除 Config */
function handleDelete(id: number) {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    state.loading = true;
    const rsp = await deleteConfigDelete({ ids: [id] });
    state.loading = false;
    if (rsp.result == 0) {
      ElMessage.success("删除成功");
      resetQuery();
    }
  });
}

/** 批量删除 */
const bantchDelete = () => {
	if(state.ids.length<=0) {
		ElMessage.warning("请勾选删除项");
		return;
	}
	ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
		confirmButtonText: "确定",
		cancelButtonText: "取消",
		type: "warning",
	}).then(async () => {
		state.loading = true;
		const rsp = await deleteConfigDelete({ ids: state.ids });
		state.loading = false;
		if(rsp.result == 0) {
			ElMessage.success("删除成功");
		}
		resetQuery();
	});
};

onMounted(() => {
  fetchPage();
});
</script>