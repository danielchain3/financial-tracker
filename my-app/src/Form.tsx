import React, { useState } from 'react';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import LocalizationProvider from '@mui/lab/LocalizationProvider';
import TextField from '@mui/material/TextField';
import DatePicker from '@mui/lab/DatePicker';
import Stack from '@mui/material/Stack';

function Form() {
    const [submitting, setSubmitting] = useState(false);
    const [day, setDay] = useState(null);

    const handleSubmit = (event: any) =>
    {
        event.preventDefault();
        setSubmitting(true);

        setTimeout(() => {
            setSubmitting(false);
        }, 3000);
    }

    const handleDayChange = (day: any) => {setDay(day);}

    return (
            <div className='wrapper'>
                <h1>Transaction</h1>
                { submitting && 
                <div>Submitting Form... </div>
                }
                <form onSubmit={handleSubmit}>
                    <fieldset>
                    <Stack spacing={3}>
                        <label>
                            <p>Account Name</p>
                            <input name='accName' />
                        </label>
                        <label>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                              <DatePicker
                                label="Date"
                                value={day}
                                onChange={handleDayChange}
                                renderInput={(params) => <TextField {...params} />}
                              />
                            </LocalizationProvider>
                        </label>
                        <label>
                            <p>Spending Category</p>
                            <input name='spendingCategory' />
                        </label>
                        <label>
                            <p>Spending Name</p>
                            <input name='spendingName' />
                        </label>
                        <label>
                            <p>Amount Spent</p>
                            <input name='amountSpent' />
                        </label>
                        <label>
                            <p>Description</p>
                            <input type='text' name='spendingDescription' />
                        </label>
                    </Stack>
                    </fieldset>

                    <button type='submit'>Submit</button>
                </form>
            </div>
            );
}

export default Form; 
