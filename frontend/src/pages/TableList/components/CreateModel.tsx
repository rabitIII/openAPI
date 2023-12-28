import { ProColumns, ProTable } from '@ant-design/pro-components';
import '@umijs/max';
import { Modal } from 'antd';

export type Props = {
  columns: ProColumns<API.InterfaceInfo>[];
  onCancel: () => void; // 点击 ‘X’或者 ‘取消’ 按钮关闭对话框
  onSubmit: (values: API.InterfaceInfo) => Promise<void>;
  open: boolean; // 对话框是否可见
  // values: Partial<API.RuleListItem>;
};

const CreateModel: React.FC<Props> = (props) => {
  const { open, columns, onCancel, onSubmit } = props;
  return (
    <Modal open={open} onCancel={() => onCancel?.()}>
      <ProTable
        type="form"
        columns={columns}
        onSubmit={async (value) => {
          onSubmit?.(value);
        }}
      />
    </Modal>
  );
};

export default CreateModel;
