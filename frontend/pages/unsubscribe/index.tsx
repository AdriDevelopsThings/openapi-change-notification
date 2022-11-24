import { useRouter } from 'next/router.js';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import errorSwal from '../../utils/errorSwal';

const unsubscribeEmail = (email: string) => {
    Swal.showLoading(null)
    fetch(process.env.NEXT_PUBLIC_API_BASE + '/unsubscribe?email=' + encodeURIComponent(email), { method: 'POST' })
        .then(r => r.json())
        .then(j => {
            if (Object.hasOwn(j, 'error')) {
                errorSwal(j.error)
            } else if (j.status == 'success') {
                Swal.fire('Success', 'We sent you an unsubscribe email. Please verify your identity by clicking on the unsubscribe link in this email.', 'success')
            } else {
                errorSwal('default')
            }
        }).catch(e => {
            errorSwal('default')
        })
}

export default function Unsubscribe() {
    const router = useRouter()
    const { email } = router.query
    useEffect(() => {
        if (!email) {
            Swal.fire({
                title: 'Unsubscribe',
                text: 'Whats your email address?',
                inputPlaceholder: 'Mail address',
                input: 'text',
                showCancelButton: true,
                allowOutsideClick: () => !Swal.isLoading(),
                preConfirm: (inputValue) => unsubscribeEmail(inputValue),
            })
        } else {
            unsubscribeEmail(email as string)
        }
    }, [email])
}