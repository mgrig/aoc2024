function day05
  dist = load_directed_graph('05_rules.txt', 99);

  for k = 1:size(dist,1)
      dist(k, k) = 0;
  endfor

  updates = load_comma_separated_struct_array('05_updates.txt');
  sum = 0;
  for r = 1:length(updates)
    line = updates(r).data;
    right_order = true;

    dist_copy = Inf(99, 99);
##    keep only rules among the numbers in 'line'
    for ir = 1:length(line)
      for ic = 1:length(line)
        dist_copy(line(ir), line(ic)) = dist(line(ir), line(ic));
      endfor
    endfor

    dist_copy = compute_fw(dist_copy);

    for c = 1:length(line)-1
     d = dist_copy(line(c), line(c+1));
##     fprintf("%d -(%d)-> ", line(c), d);
     if isinf(d) # || d == 0
        right_order = false;
        break
      endif
    endfor
##    fprintf("%d\n", line(c+1))

    if right_order
      mid_value = line((1+length(line))/2);
##      fprintf("%d %d\n", r, mid_value);
      sum = sum + mid_value;
    endif
  endfor

  disp(sum)
end

function dist = compute_fw(dist)
## Use Floyd-Warshall algorithm to calculate the shortest path lengths between every pair of vertices in a graph
  for k = 1:size(dist,1)
    dist = min (dist, dist(:,k) + dist(k,:));
  endfor
end

function adj_matrix = load_directed_graph(filename, n)

    % Check if the file exists
    if ~exist(filename, 'file')
        error('File not found: %s', filename);
    end

    % Open the file for reading
    fid = fopen(filename, 'r');
    if fid == -1
        error('Cannot open file: %s', filename);
    end

    % Read the entire file using textscan
    % Each line is expected to have two integers separated by '|'
    data = textscan(fid, '%d|%d', 'Delimiter', '|', 'CollectOutput', true, ...
                   'TreatAsEmpty', {'', 'Inf'}, 'CommentStyle', '#');
    fclose(fid);

    % Extract the edges from the scanned data
    edges = data{1}; % Nx2 matrix where each row is [from, to]

    % Check if any valid edges were found
    if isempty(edges)
        warning('No valid edges found in the file.');
        adj_matrix = [];
        return;
    end

    % Extract 'from' and 'to' nodes
    from_nodes = edges(:, 1);
    to_nodes = edges(:, 2);

    % Determine the maximum node number to size the adjacency matrix
##    max_from = max(from_nodes);
##    max_to = max(to_nodes);
##    max_node = max(max_from, max_to);
    max_node = n;

    % Initialize the adjacency matrix with Inf
    adj_matrix = Inf(max_node, max_node);

    % Convert 'from' and 'to' node pairs to linear indices
    % and set the corresponding entries in the adjacency matrix to 1
    indices = sub2ind(size(adj_matrix), from_nodes, to_nodes);
    adj_matrix(indices) = 1;

##    % Optional: Display a message indicating successful loading
##    fprintf('Loaded directed graph from "%s".\n', filename);
##    fprintf('Number of nodes: %d\n', max_node);
##    fprintf('Number of edges: %d\n', length(indices));
end

function matrix_struct = load_comma_separated_struct_array(filename)
    %LOAD_COMMA_SEPARATED_STRUCT_ARRAY Loads a variable-length matrix from a file.  updates = load_comma_separated_struct_array('05_updates.txt');

    %   matrix_struct = LOAD_COMMA_SEPARATED_STRUCT_ARRAY(filename) reads a file where each
    %   line contains a comma-separated list of integers with no spaces. Each line can have
    %   a different number of integers. It returns a struct array where each struct corresponds
    %   to a row in the file, containing a field 'data' with the row vector of integers.
    %
    %   Input:
    %       filename - String representing the path to the input file.
    %
    %   Output:
    %       matrix_struct - 1xM struct array where each struct has a field 'data'
    %                       containing a 1xN_i integer row vector.
    %
    %   Example:
    %       If the file "data_variable.csv" contains:
    %           1,2,3
    %           4,5
    %           6,7,8,9
    %       Then, calling:
    %           mat_struct = load_comma_separated_struct_array('data_variable.csv');
    %       Will return:
    %           mat_struct(1).data = [1 2 3]
    %           mat_struct(2).data = [4 5]
    %           mat_struct(3).data = [6 7 8 9]

    % Check if the file exists
    if ~exist(filename, 'file')
        error('File not found: %s', filename);
    end

    % Open the file for reading
    fid = fopen(filename, 'r');
    if fid == -1
        error('Cannot open file: %s', filename);
    end

    % Initialize an empty struct array
    matrix_struct = struct('data', {});

    % Read the file line by line
    line_num = 0;
    while ~feof(fid)
        line = fgetl(fid);
        line_num = line_num + 1;

        % Trim whitespace from the line
        line = strtrim(line);

        % Skip empty lines
        if isempty(line)
            continue;
        endif

        % Optionally, skip comment lines (starting with #)
        if startsWith(line, '#')
            continue;
        endif

        % Split the line by commas
        parts = strsplit(line, ',');

        % Convert each part to an integer
        try
            numbers = cellfun(@str2num, parts, 'UniformOutput', true);
        catch
            fclose(fid);
            error('Non-integer value found on line %d: "%s".', line_num, line);
        end_try_catch

        % Check for successful conversion
        if any(isnan(numbers))
            fclose(fid);
            error('Invalid integer found on line %d: "%s".', line_num, line);
        endif

        % Append to the struct array
        matrix_struct(end + 1).data = numbers;
    endwhile

    fclose(fid);

    % Check if any data was read
    if isempty(matrix_struct)
        warning('No data found in the file: %s.', filename);
    endif
end

