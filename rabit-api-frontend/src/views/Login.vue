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
  
<script setup>
import { register } from '@/apis/auth';
import { useAxios } from '@/utils/request';
import { Message } from '@arco-design/web-vue';
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const form = reactive({
    userAccount: '',
    userPassword: '',
});
async function handleSubmit() {
    useAxios.post('/user/login', {
        user_account: form.userAccount,
        user_password: form.userPassword,
    }).then(res => {
        if (res.code === 0) {
            // setting token
            if (res.data.user.role === 1) {
                router.push("/admin")
            } else if (res.data.user.role === 0) {
                router.push("/user")
            }
            Message.success('登录成功')
        }
    })
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