import { useState } from "react";

type TaskProps = {
    name: string;
    isCompleted: boolean;
}

function Task({name, isCompleted} : TaskProps) {

    const [completed, setCompleted] = useState<boolean>(isCompleted)

    const handleToggle = () => {
        setCompleted(prev => !prev);
    }

    return (
        <div className="m-10 py-2 px-4 flex border-emerald-600 border-1 rounded-2xl">
            <p className="m-2 w-48 text-left">{name}</p>
            <input onChange={handleToggle} type="checkbox" checked={completed}></input>
        </div>
    )
}

export default Task;