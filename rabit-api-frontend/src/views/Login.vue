<template>
    <a-watermark content="rabitApi">
        
    <div class="login_page">
        <div style="margin: 10px; padding: 10px; border: 1px solid gray; border-radius: 10px;">
            <h1 style="text-align: center">登录页面</h1>
            <a-form :model="form" :style="{ width: '600px' }" @submit="handleSubmit">
                <a-form-item field="name" tooltip="Please enter username" label="Username">
                    <a-input
                    v-model="form.userAccount"
                    placeholder="please enter your userAccount..."
                    />
                </a-form-item>
                <a-form-item field="post" label="Post">
                    <a-input v-model="form.userPassword" placeholder="please enter your passwrod..." />
                </a-form-item>
                <a-form-item>
                    <a-button html-type="submit">Submit</a-button>
                </a-form-item>
            </a-form>
        </div>
    </div>
    </a-watermark>
</template>
  
<script>
import { reactive } from 'vue';

export default {
setup() {
    const form = reactive({
        userAccount: '',
        userPassword: '',
    });
    const handleSubmit = (data) => {
    console.log(data);
    };

    return {
    form,
    handleSubmit,
    };
},
};
</script>
  
<style>
.login_page {
    display: flex;
    width: 100%;
    height: 100vh;
    justify-content: center;
    align-items: center;
    background-color: var(--color-bg-white);
}
</style>