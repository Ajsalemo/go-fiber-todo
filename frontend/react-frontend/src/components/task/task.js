import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircleCheck } from '@fortawesome/free-regular-svg-icons'
import { faCircleCheck as faSolidCircleCheck } from '@fortawesome/free-solid-svg-icons'

export const Task = ({ taskName, completed }) => (
    <div className="bg-slate-800 w-4/5 min-h-16 rounded-md p-4 mt-8 mx-8 flex shadow-2xl align-center">
        <div>
            <FontAwesomeIcon icon={completed ? faSolidCircleCheck : faCircleCheck} className="text-white mr-2" />        
        </div>
        <span className="break-all text-white">
            {taskName}
        </span>
    </div>
)