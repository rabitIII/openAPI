<template>
    <a-layout class="layout-demo">
        <a-layout-sider
        hide-trigger
        collapsible
        :collapsed="collapsed"
        >
        <div class="logo">RabitApi</div>
        <a-menu
            :defaultOpenKeys="['1']"
            :defaultSelectedKeys="['dashboard']"
            :style="{ width: '100%' }"
            @menuItemClick="onClickMenuItem"
        >
            <a-menu-item key="dashboard">
                <icon-dashboard />
                dashboard
            </a-menu-item>
            <a-sub-menu key="1">
                <template #title>
                    <span><IconCalendar />管理页</span>
                </template>
                <a-menu-item key="interface">
                    <icon-branch />
                    接口管理
                </a-menu-item>
            </a-sub-menu>
            
        </a-menu>
        </a-layout-sider>
        <a-layout>
        <div class="rabit_header">
            <a-button class="side-button" shape="round" @click="onCollapse">
                <IconCaretRight v-if="collapsed" />
                <IconCaretLeft v-else />
            </a-button>
            <div class="rabit_header_menu">
                <a-dropdown :popup-max-height="false">
                Welcome，
                    <a-button> Admin / User </a-button>
                    <template #content>
                        <a-doption>option 1</a-doption>
                        <a-doption>
                            <a-button @click="handleLogout" type="text" status="danger">退出</a-button>
                        </a-doption>
                    </template>
                  </a-dropdown>
            </div>
        </div>
        <a-layout style="padding: 0 24px;">
            <a-breadcrumb :style="{ margin: '16px 0' }">
            <a-breadcrumb-item style="font-size: 1.5rem">{{ $route.meta.title }}</a-breadcrumb-item>
            </a-breadcrumb>
            <a-layout-content>
                <RouterView />
            </a-layout-content>
            <a-layout-footer>Footer</a-layout-footer>
        </a-layout>
        </a-layout>
    </a-layout>
</template>

<script>
import { defineComponent, ref } from 'vue';
import { Message, Modal } from '@arco-design/web-vue';
import {
IconCaretRight,
IconCaretLeft,
IconHome,
IconCalendar,
} from '@arco-design/web-vue/es/icon';
import {useRoute, useRouter} from "vue-router";
import Logout from "@/views/Admin/Logout.vue";


export default defineComponent({
components: {
    Logout,
    IconCaretRight,
    IconCaretLeft,
    IconHome,
    IconCalendar,
},
setup() {
    const collapsed = ref(false);
    const router = useRouter()
    const route = useRoute()
    const onCollapse = () => {
    collapsed.value = !collapsed.value;
    };
    
    const handleLogout = () => {
        Modal.error({
          title: '确认退出？',
          content: '操作不可逆！',
            onOk:e => {
                router.push("/")
                Message.info({ content: `你退出了管理中心`, showIcon: true });
            }
        });
    };
    
    const onClickMenuItem = (key) => {
        router.push({ name: key })
    }

    return {
        collapsed,
        onCollapse,
        onClickMenuItem,
        
        handleLogout,
    };
},
});

</script>

<style scoped>
.layout-demo {
    height: 100vh;
    background: var(--color-fill-2);
/* border: 1px solid var(--color-border); */
}
.layout-demo :deep(.arco-layout-sider) .logo {
    height: 32px;
    margin: 12px 8px;
    font-size: 1.5rem;
    background: rgba(255, 255, 255, 0.2);
}
.layout-demo :deep(.arco-layout-sider-light) .logo{
    background: var(--color-fill-2);
}
.layout-demo :deep(.arco-layout-header)  {
    height: 64px;
    line-height: 64px;
    background: var(--color-bg-3);
}

.rabit_header {
    display: flex;
    height: 64px;
    background: var(--color-bg-3);
    justify-content: space-between;
    align-items: center;
    .side-button {
        margin-left: 20px;
    }
    .rabit_header_menu {
        margin-right: 24px;
    }
}

.layout-demo :deep(.arco-layout-footer) {
height: 48px;
color: var(--color-text-2);
font-weight: 400;
font-size: 14px;
line-height: 48px;
}
.layout-demo :deep(.arco-layout-content) {
color: var(--color-text-2);
font-weight: 400;
font-size: 14px;
/*background: var(--color-bg-3);*/
}
.layout-demo :deep(.arco-layout-footer),
.layout-demo :deep(.arco-layout-content)  {
display: flex;
flex-direction: column;
/*justify-content: center;*/
/*color: var(--color-white);*/
font-size: 16px;
font-stretch: condensed;
text-align: center;
}
</style>
  