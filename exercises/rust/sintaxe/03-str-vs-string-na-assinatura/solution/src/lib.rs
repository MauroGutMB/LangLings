// `&str` aceita tanto um literal quanto uma `&String` (a coerção de deref
// cuida da segunda forma) — por isso uma função que só LÊ texto deveria
// pedir `&str`, nunca `&String`.
pub fn contem_arroba(s: &str) -> bool {
    s.contains('@')
}

pub fn validar_emails(emails: &[&str]) -> Vec<bool> {
    emails.iter().map(|e| contem_arroba(e)).collect()
}
