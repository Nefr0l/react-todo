import Task from "../Task/Task";


function TaskList() {
    return (
        <>
            <ul className="m-16">
                <li><Task name='xd' isCompleted></Task></li>
                <li><Task name='2' isCompleted={false}></Task></li>
            </ul>
        </>
    );
}

export default TaskList;