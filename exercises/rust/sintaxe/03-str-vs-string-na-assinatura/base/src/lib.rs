// TODO: isto não compila. `validar_emails` chama `contem_arroba` com
// literais (`&str`), e `&str` não vira `&String` sozinho — só o contrário.
pub fn contem_arroba(s: &String) -> bool {
    s.contains('@')
}

pub fn validar_emails(emails: &[&str]) -> Vec<bool> {
    emails.iter().map(|e| contem_arroba(e)).collect()
}
