import React from "react";

interface Task {
  id: string;
  title: string;
  description: string;
  status: string;
}

interface TaskDashboardProps {
  tasks: Task[];
}

const TaskDashboard: React.FC<TaskDashboardProps> = ({ tasks }) => {
  return (
    <div className="space-y-4">
      {tasks.map((task) => (
        <div key={task.id} className="p-4 border rounded shadow">
          <h2 className="text-xl font-semibold">{task.title}</h2>
          <p>{task.description}</p>
          <p>Status: {task.status}</p>
        </div>
      ))}
    </div>
  );
};

export default TaskDashboard;
