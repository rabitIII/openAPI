import { ProColumns, ProTable } from '@ant-design/pro-components';
import '@umijs/max';
import { Modal } from 'antd';

export type Props = {
  values: API.InterfaceInfo;
  columns: ProColumns<API.InterfaceInfo>[];
  onCancel: () => void; // 点击 ‘X’或者 ‘取消’ 按钮关闭对话框
  onSubmit: (values: API.InterfaceInfo) => Promise<void>;
  visible: boolean; // 对话框是否可见
};

const UpdateModel: React.FC<Props> = (props) => {
  const { values, visible, columns, onCancel, onSubmit } = props;
  return (
    <Modal open={visible} onCancel={() => onCancel?.()}>
      <ProTable
        type="form"
        columns={columns}
        form={{
          initialValues: values,
        }}
        onSubmit={async (value) => {
          onSubmit?.(value);
        }}
      />
    </Modal>
  );
};

export default UpdateModel;
