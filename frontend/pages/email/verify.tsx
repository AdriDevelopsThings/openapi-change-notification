import { useRouter } from 'next/router.js';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import errorSwal from '../../utils/errorSwal';

export default function Verify() {
    const router = useRouter()
    const { code } = router.query

    useEffect(() => {
        if (code) {
            fetch(process.env.NEXT_PUBLIC_API_BASE + '/email/verify/' + code, { method: 'POST' })
                .then(r => r.json())
                .then(r => {
                    if (Object.hasOwn(r, 'error')) {
                        errorSwal(r.error)
                    } else if (r.status == 'success') {
                        Swal.fire({
                            title: 'Success', 
                            text: 'Your email address is now verified. You will now get OpenAPI deprecation change notifications.',
                            icon: 'success',
                            willClose: () => { window.location.href = '/' }
                        })
                    } else {
                        errorSwal('default')
                    }
                }).catch(e => {
                    console.error(e)
                    errorSwal('default')
                })
        } else {
            Swal.fire({
                title: 'Verify email',
                text: 'Verify your email address by clicking on the link we sent to you. Don\'t you have a link or your link expired? Just send a new or your old subscription on the home site again. You will get a new verification email.',
                icon: 'warning',
                willClose: () => { window.location.href = '/' },
            })
        }
    }, [code])
}