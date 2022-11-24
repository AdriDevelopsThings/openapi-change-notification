import { useRouter } from 'next/router.js';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import errorSwal from '../../utils/errorSwal';

export default function Verify() {
    const router = useRouter()
    const { code } = router.query
    useEffect(() => {
        if (code) {
            fetch(process.env.NEXT_PUBLIC_API_BASE + '/unsubscribe/verify/' + code, { method: 'POST' })
                .then(r => r.json())
                .then(r => {
                    if (Object.hasOwn(r, 'error')) {
                        errorSwal(r.error)
                    } else if (r.status == 'success') {
                        Swal.fire({
                            title: 'Success', 
                            text: 'You unsubscribed from all OpenAPI deprecation notifcations.',
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
                title: 'Unsubscribe',
                text: 'Unsubscribe from all OpenAPI deprecation notifications by submitting the form on /unsubscribe.',
                icon: 'warning',
    
                willClose: () => { window.location.href = '/unsubscribe' },
            })
        }
    }, [code])
}