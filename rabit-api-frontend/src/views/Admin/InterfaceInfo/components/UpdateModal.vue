<template>
  <a-button type="primary" @click="handleClick">修改</a-button>
  <a-modal v-model:visible="visible" title="Modal Form" @cancel="handleCancel" @before-ok="handleBeforeOk">
    <a-form :model="form">
      <a-form-item field="name" tooltip="Please enter username" label="接口名称">
        <a-input
          v-model="form.name"
          placeholder="please enter your username..."
        />
      </a-form-item>
      <a-form-item field="description" label="描述">
        <a-input v-model="form.description" placeholder="描述信息..." />
      </a-form-item>
      <a-form-item field="method" label="请求方法">
        <a-select v-model="form.method" placeholder="Please select ..." allow-clear>
          <a-option value="section one">GET</a-option>
          <a-option value="section two">POST</a-option>
          <a-option value="section three">PUT</a-option>
          <a-option value="section three">DELETE</a-option>
        </a-select>
      </a-form-item>
      <a-form-item field="url" label="url">
        <a-input v-model="form.url" placeholder="e.g. www.baidu.com" />
      </a-form-item>
      <a-form-item field="drequestParams" label="请求参数">
        <a-input v-model="form.requestParams" placeholder="" />
      </a-form-item>
      <a-form-item field="requestHeader" label="请求头">
        <a-input v-model="form.requestHeader" placeholder="请求头信息..." />
      </a-form-item>
      <a-form-item field="responseHeader" label="响应头">
        <a-input v-model="form.responseHeader" placeholder="响应头信息..." />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script>
import { reactive, ref } from 'vue';

export default {
  setup() {
    const form = reactive({
      name: '',
      description: '',
      method: '',
      url: '',
      requestParams: '',
      requestHeader: '',
      responseHeader: '',
      status: false,
    });
    const visible = ref(false);
    const handleClick = () => {
      visible.value = true;
    };
    const handleBeforeOk = (done) => {
      console.log(form)
      window.setTimeout(() => {
        done()
        // prevent close
        // done(false)
      }, 3000)
    };
    const handleCancel = () => {
      visible.value = false;
    }

    return {
      form,
      visible,
      handleClick,
      handleBeforeOk,
      handleCancel
    }
  },
};
</script>
