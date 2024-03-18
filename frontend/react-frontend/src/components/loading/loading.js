import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSpinner } from '@fortawesome/free-solid-svg-icons'

export const Loading = () => {
    return (
        <div className="h-full flex justify-center items-center">
            <FontAwesomeIcon icon={faSpinner} className="text-white mr-2 animate-spin" />
        </div>
    )
}