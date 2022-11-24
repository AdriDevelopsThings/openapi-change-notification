import Swal from 'sweetalert2';

const translation: { [key: string]: { title: string, message: string, redirect?: string}} = {
    default: { title: 'Unknown error', message: 'A unknown error has occured. Please check your inputs or contact the support.'},
    bad_request_error: { title: 'Wrong input', message: 'Please recheck your input.'},
    open_api_fetching_error: { title: 'OpenAPI fetching error', message: 'There was an error while fetching the openapi configuration by the server. Please recheck your input.'},
    path_could_not_be_found_error: { title: 'Path not found', message: 'The specified path could not be found in the openapi configuration file.'},
    path_method_could_not_be_found_error: { title: 'Path method missmatch', message: 'The path on the openapi configuration doesn\' implement this method. Please recheck.'},
    subscription_not_found_error: { title: 'Subscription not found', message: 'The subscription wasn\'t found.'},
    email_verification_code_error: { title: 'Email verification code invalid', message: 'The email verification code is invalid or expired. Please get a new one by resubscribing.', redirect: '/'},
    unsubscribe_verification_code_error: { title: 'Unsubscribe verification code invalid', message: 'The unsubscribe verification code is invalid or expired. Please get a new one by unsubscribing.', redirect: '/unsubscribe'}
}

export default function errorSwal(name: string) {
    const t = translation[name] || translation['default']
    Swal.fire({
        title: t.title,
        text: t.message,
        icon: 'error',
        willClose: () => {
            if (t.redirect) {
                window.location.href = t.redirect
            }
        }
    })
}