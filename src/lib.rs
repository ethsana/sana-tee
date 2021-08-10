// SPDX-License-Identifier: Apache-2.0
// #![deny(clippy::all)]
// #![deny(missing_docs)]

mod error;

use std::ptr;
use std::io::{Error, ErrorKind};

use error::{Contextual, Result};

use ::sev::firmware::{Firmware};


#[no_mangle]
pub extern "C" fn sev_device_id(result_string : * mut u8 ) {
    let id = derive_id().unwrap();

    unsafe{
        ptr::copy_nonoverlapping(id.as_ptr(), result_string, id.len());
    }
}

fn firmware() -> Result<Firmware> {
    Firmware::open().context("unable to open /dev/sev")
}


fn derive_id() -> Result<String> {
    let id = firmware()?
        .get_identifier()
        .map_err(|e| Error::new(ErrorKind::Other, format!("{:?}", e)))
        .context("error fetching identifier")?;

    Ok(format!("{}",id))
}

