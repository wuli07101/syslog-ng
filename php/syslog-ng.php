<?PHP
while($line = fgets(STDIN)) {
    $line_parts = array();
    $line_parts = preg_split('/[\t]/',$line);
	file_put_contents("./test.txt",var_export($line_parts,TRUE),FILE_APPEND);
}
