import { useEffect, useState } from "react";
import { RedoOutlined } from "@ant-design/icons";
import { Table, Button, Form, Input, Col, Row, Divider } from "antd";
import CreateModal from "./components/CreateModal";
import {
  addInterfaceInfoPost,
  listInterfaceInfoGet,
  updateInterfaceInfoPut,
} from "../../../api/interfaceInfoApi";
import SearchList from "../../../components/SearchList";
import UpdateModal from "./components/UpdateModal";

const { Search } = Input;
const data = [
  {
    id: "1",
    title: "Jim Green",
    desc: "这是一个接口",
    method: "POST",
    url: "https://www.baidu.com",
    requestHeader: "xxxxxx",
    responseHeader: "json字符串",
    status: ["关闭"],
  },
  {
    id: "2",
    title: "Jim Green",
    desc: "这是一个接口",
    method: "POST",
    url: "https://www.baidu.com",
    requestHeader: "xxxxxx",
    responseHeader: "json字符串",
    status: ["关闭"],
  },
  {
    id: "3",
    title: "Jim Green",
    desc: "这是一个接口",
    method: "POST",
    url: "https://www.baidu.com",
    requestHeader: "xxxxxx",
    responseHeader: "json字符串",
    status: ["关闭"],
  },
];

// rowSelection object indicates the need for row selection
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(
      `selectedRowKeys: ${selectedRowKeys}`,
      "selectedRows: ",
      selectedRows
    );
  },
  getCheckboxProps: (record) => ({
    disabled: record.name === "Disabled User",
    // Column configuration not to be checked
    name: record.name,
  }),
};

const AdminInterfaceInfo = () => {
  // 接口状态管理
  const [createModalVisible, handleModalVisible] = useState(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState(false);
  //   //   const [showDetail, setShowDetail] = useState(false);
  const [titlesearch, setTitleSearch] = useState(null); // 搜索框的值
  const [currentRow, setCurrentRow] = useState();
  const [selectionType, setSelectionType] = useState("checkbox");
  const [interfaceInfoList, setInterfaceInfoList] = useState([]); // table表数据
  // const [clearSearch, setClearOnSearch] = useState(false);

  // const [form] = Form.useForm();

  const loadApiList = () => {
    listInterfaceInfoGet(titlesearch).then(({ data }) => {
      data = data.list.map((record) => {
        return {
          ...record,
          key: record.ID,
        };
      });
      setInterfaceInfoList(data);
    });
  };

  useEffect(() => {
    loadApiList();
  }, []);

  // 模糊查询
  const onSearch = (values) => {
    // setClearOnSearch(true);
    console.log("查询", values);
    loadApiList(titlesearch);
  };

  //   const style = {
  //     background: "#0092ff",
  //     padding: "8px 0",
  //   };
  /**
   * 添加接口(默认不对外开放)
   *
   */
  const handleAdd = async (values) => {
    addInterfaceInfoPost(values);
  };

  /**
   * 更新接口
   *
   */
  const handleUpdate = async (values) => {
    if (!currentRow){
      return;
    }
    try {
      await updateInterfaceInfoPut({
        ID: currentRow.ID,
        ...values});
      return true;
    }catch (error) {
      console.log("错误：",error.message)
      return false;
    }

  };

  /**
   * 删除接口(从表格中删除,数据库不删除)
   *
   */
  const handleRemove = async () => {};

  /**
   * 发布接口
   *
   */
  const handleOnline = async () => {};

  /**
   * 下线接口
   *
   */
  const handleOffline = async () => {};

  /**
   * 接口管理列表的参数类型
   */
  const columns = [
    {
      title: "id",
      dataIndex: "ID",
      key: "id",
    },
    {
      title: "接口名称",
      dataIndex: "title",
      key: "title",
    },
    {
      title: "描述",
      dataIndex: "description",
      key: "desc",
    },
    {
      title: "请求方法",
      dataIndex: "method",
      key: "method",
    },
    {
      title: "url",
      dataIndex: "url",
      key: "url",
      render: (text) => <a>{text}</a>,
    },
    {
      title: "请求头",
      dataIndex: "requestHeader",
      key: "requestHeader",
    },
    {
      title: "响应头",
      dataIndex: "responseHeader",
      key: "responseHeader",
      //   render: (text) => <a>{text}</a>,
    },
    {
      title: "状态",
      dataIndex: "status",
      key: "status",
    },
    {
      title: "操作",
      dataIndex: "option",
      key: "option",
      render: (_, record) => [
        <a
          key="config"
          onClick={() => {
            handleUpdateModalVisible(true);
            setCurrentRow(record);
            console.log("record:", record);
            // console.log("currentRow:",currentRow)
          }}
        >
          修改
        </a>,
        record.status === 0 ? (
          <a
            key="online"
            onClick={() => {
              handleOnline(record);
            }}
          >
            发布
          </a>
        ) : null,
        record.status === 1 ? (
          <Button
            type="text"
            key="logout"
            danger
            onClick={() => {
              handleOffline(record);
            }}
          >
            下线
          </Button>
        ) : null,
        <Button
          type="text"
          key="delete"
          danger
          onClick={async () => {
            await handleRemove(record);
          }}
        >
          删除
        </Button>,
      ],
      // (
      //   <Space size="middle">
      //     <a>Invite {record.name}</a>
      //     <a>Delete</a>
      //   </Space>
      // ),
    },
  ];
  return (
    <div>
      <Search
        id="title"
        size="large"
        style={{
          width: "200px",
        }}
        value={titlesearch}
        onChange={(e) => {
          setTitleSearch(e.target.value);
        }}
        placeholder="input search apiName"
        onSearch={onSearch}
        enterButton
      />
      {/* <SearchList /> */}

      {/* <Radio.Group
        onChange={({ target: { value } }) => {
          setSelectionType(value);
        }}
        value={selectionType}
      >
        <Button
          onClick={async () => {
            await handleRemove(setSelectionType);
          }}
        >
          批量删除
        </Button>
      </Radio.Group> */}

      <Divider />
      <div
        style={{
          maxWidth: "100vw",
          display: "flex",
          justifyContent: "space-between",
          marginBottom: "1rem",
        }}
      >
        <div style={{ fontSize: "1rem" }}>查询表格</div>
        <div>
          <Button
            type="primary"
            onClick={() => {
              handleModalVisible(true);
            }}
          >
            新建
          </Button>
        </div>
      </div>
      <Table
        //   多选框
        rowSelection={{
          type: selectionType,
          ...rowSelection,
        }}
        columns={columns}
        dataSource={interfaceInfoList}
      />
      <CreateModal
        open={createModalVisible}
        onSubmit={(values) => {
          handleAdd(values);
        }}
        onCancel={() => {
          handleModalVisible(false);
        }}
      />
      <UpdateModal
        visible={updateModalVisible}
        values={currentRow || {}}
        onSubmit={ (values) => {
          console.log("father:", values)
          handleUpdate(values);
        }}
        onCancel={() => {
          handleUpdateModalVisible(false);
        }}
      />
    </div>
  );
};
export default AdminInterfaceInfo;
