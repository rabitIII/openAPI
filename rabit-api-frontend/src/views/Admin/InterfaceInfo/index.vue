<template>
<div class="rabit_table">
    <div class="rabit_head">
        <div class="action_create">
            <create-modal  />
        </div>
        <div class="action_select">
            <a-select :style="{width:'120px'}" placeholder="method">
                <a-option>GET</a-option>
                <a-option>POST</a-option>
                <a-option>PUSH</a-option>
                <a-option>DELETE</a-option>
            </a-select>
        </div>
        
        <div class="action_search">
            <a-input-search :style="{width:'320px'}" placeholder="Please enter something" search-button />
        </div>
        <div class="action_flush"><icon-refresh /></div>
    </div>
    
    <div class="rabit_table_source">
        <a-table :columns="columns" :data="data" :row-selection="rowSelection" :pagination="false">
            <template #action="{ record }">
                <UpdateModal />
                <a-popconfirm content="是否要进行删除操作?" type="warning">
                  <a-button status="danger" type="primary">删除</a-button>
                </a-popconfirm>
            </template>
        </a-table>
    </div>
    <div class="rabit_table_page">
        <a-pagination :total="50" show-total show-jumper/>
    </div>
</div>

    
</template>

<script>
import {reactive} from 'vue';
import CreateModal from "@/views/Admin/InterfaceInfo/components/CreateModal.vue";
import UpdateModal from "@/views/Admin/InterfaceInfo/components/UpdateModal.vue";

export default {
    components: {UpdateModal, CreateModal},
    setup() {
        const columns = [
            {
                title: '接口名称',
                dataIndex: 'name',
                fixed: 'left',
                width: 140
            },
            {
                title: '描述',
                dataIndex: 'description',
            },
            {
                title: '请求方法',
                dataIndex: 'method',
            },
            {
                title: 'url',
                dataIndex: 'url',
            },
            {
                title: '请求参数',
                dataIndex: 'requestParams',
            },
            {
                title: '请求头',
                dataIndex: 'requestHeader',
            },
            {
                title: '响应头',
                dataIndex: 'responseHeader',
            },
            {
                title: '状态',
                dataIndex: 'status',
            },
            {
                title: '操作',
                slotName: 'action',
            },
        ];
        const data = reactive(Array(5).fill(null).map((_, index) => ({
            key: String(index),
            name: `User ${index + 1}`,
            description: '32 Park Road, London',
            method: `POST`,
            url: '',
            requestParams: '',
            requestHeader: '',
            responseHeader: '',
            status: '',
        })));
        const rowSelection = {
            type: 'checkbox',
            showCheckedAll: true
        };
    
        return {
            columns,
            data,
            rowSelection,
        }
    },
}
</script>
  
<style>
.rabit_table {
    background-color: var(--color-bg-1);
    border-radius: 5px;
    .rabit_head {
        padding: 20px 20px 10px 20px;
        border-bottom: 1px solid var(--bg);
        display: flex;
        align-items: center;
        position: relative;
        
        >div {
            margin-right: 10px;
        }
        
        .action_flush {
            position: absolute;
            right: 10px;
        }
    }
    .rabit_table_source {
        padding: 10px 20px 20px 20px;
    }
    .rabit_table_page {
        display: flex;
        justify-content: center;
        margin: 10px;
    }
}
</style>